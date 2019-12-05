FROM calico/node:v3.9.1

ADD routing/routing /bin/gobgp-vpp-agent
ADD config/service/gobgp-vpp-agent /etc/service/available/gobgp-vpp-agent

RUN sed -i.orig '/^case "\$CALICO_NETWORKING_BACKEND" in/a \\t"vpp" )\n\
\tcp -a /etc/service/available/gobgp-vpp-agent /etc/service/enabled/\n\
\tsh -c '\''for file in `find /etc/calico/confd/conf.d/ -not -name '\''tunl-ip.toml'\'' -type f`; do rm $file; done'\''\n\
\tcp -a /etc/service/available/confd /etc/service/enabled/\n\
\t;;\n' /etc/rc.local


