FROM alpine:latest
MAINTAINER "louisehong <louisehong4168@gmail.com>"
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
        && apk update \
        && apk upgrade \
        && apk add --no-cache bash  \
                       curl

ENTRYPOINT ["/entrypoint.sh"]

COPY .dist/mycli_linux_amd64/mycli /bin/mycli
COPY scripts/entrypoint.sh /entrypoint.sh

RUN  chmod +x /entrypoint.sh
