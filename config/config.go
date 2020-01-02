// Copyright (C) 2019 Cisco Systems Inc.
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

package config

const (
	DataInterfaceSwIfIndex = uint32(1) // Assumption: the VPP config ensures this is true
	CNIServerSocket        = "/var/run/calico/cni-server.sock"
	VppAPISocket           = "/var/run/vpp/vpp-api.sock"
)

var (
	VppSideMacAddress       = [6]byte{2, 0, 0, 0, 0, 2}
	ContainerSideMacAddress = [6]byte{2, 0, 0, 0, 0, 1}
)