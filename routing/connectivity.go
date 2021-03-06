// Copyright (C) 2019 Cisco Systems Inc.
// Copyright (C) 2016-2017 Nippon Telegraph and Telephone Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package routing

import (
	"net"

	"github.com/calico-vpp/calico-vpp/config"
	"github.com/calico-vpp/vpplink"
	"github.com/pkg/errors"
	calicov3 "github.com/projectcalico/libcalico-go/lib/apis/v3"
)

type connectivityProvider interface {
	addConnectivity(dst net.IPNet, destNode net.IP) error
	delConnectivity(dst net.IPNet, destNode net.IP) error
}

func (s *Server) getNodeIP(isv6 bool) net.IP {
	if isv6 {
		return s.ipv6
	} else {
		return s.ipv4
	}
}

func (s *Server) getNodeIPNet(isv6 bool) *net.IPNet {
	if isv6 {
		return s.ipv6Net
	} else {
		return s.ipv4Net
	}
}

func (s *Server) needIpipTunnel(dst net.IPNet, otherNodeIP net.IP) (ipip bool, err error) {
	ipPool := s.ipam.match(dst)
	if ipPool == nil {
		return false, nil
	}
	if ipPool.Spec.IPIPMode == calicov3.IPIPModeNever {
		return false, nil
	}
	ipNet := s.getNodeIPNet(vpplink.IsIP6(dst.IP))
	if ipPool.Spec.IPIPMode == calicov3.IPIPModeCrossSubnet && !isCrossSubnet(otherNodeIP, *ipNet) {
		return false, nil
	}

	return true, nil
}

func (s *Server) updateIPConnectivity(dst net.IPNet, otherNodeIP net.IP, IsWithdraw bool) error {
	var provider connectivityProvider = s.flat

	ipip, err := s.needIpipTunnel(dst, otherNodeIP)
	if err != nil {
		return errors.Wrapf(err, "error checking for ipip tunnel")
	}
	if ipip && config.EnableIPSec {
		provider = s.ipsec
	} else if ipip {
		provider = s.ipip
	}

	if IsWithdraw {
		return provider.delConnectivity(dst, otherNodeIP)
	} else {
		return provider.addConnectivity(dst, otherNodeIP)
	}
}
