FROM fedora:31
WORKDIR /

RUN dnf -y update
RUN dnf -y install bind-utils jq net-tools procps iproute

ADD stupid-curl /usr/bin/stupid-curl
RUN chmod +x /usr/bin/stupid-curl

ADD stupid-http /usr/sbin/stupid-http
RUN chmod 755 /usr/sbin/stupid-http

ADD entrypoint.sh /
RUN chmod 755 /entrypoint.sh

CMD ["/entrypoint.sh"]