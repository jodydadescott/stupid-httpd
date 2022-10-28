FROM alpine:latest
WORKDIR /

RUN apk update

RUN apk add bash curl bind-tools

ADD stupid-httpd /usr/sbin
RUN chmod 755 /usr/sbin/stupid-httpd

CMD ["/usr/sbin/stupid-httpd"]