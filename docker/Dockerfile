FROM alpine:latest
WORKDIR /
MAINTAINER jodyscott

RUN apk update

RUN apk add bash curl bind-tools

RUN curl -L -o /usr/sbin/stupid-http https://github.com/jodydadescott/stupid-httpd/releases/download/v1-linux/stupid-http
RUN chmod 755 /usr/sbin/stupid-http

ADD entrypoint.sh /
RUN chmod 755 /entrypoint.sh

CMD ["/entrypoint.sh"]
