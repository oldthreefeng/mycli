FROM alpine:latest
MAINTAINER "louisehong <louisehong4168@gmail.com>"
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
        && apk update \
        && apk upgrade \
        && apk add --no-cache bash jq bash-completion\
                       curl wget

ENTRYPOINT ["/entrypoint.sh"]

COPY scripts/entrypoint.sh /entrypoint.sh
COPY mycli /bin/mycli

RUN  chmod +x /entrypoint.sh