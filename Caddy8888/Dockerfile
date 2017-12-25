FROM alpine

RUN apk update && apk upgrade && \
    apk add curl && \
    rm -rf /var/cache/apk/*

RUN mkdir -p /tmp/caddy \
 && curl -sL -o /tmp/caddy/caddy_linux_amd64.tar.gz "https://caddyserver.com/download/linux/amd64" \
 && tar -zxf /tmp/caddy/caddy_linux_amd64.tar.gz -C /tmp/caddy \
 && mv /tmp/caddy/caddy /usr/bin/ \
 && chmod +x /usr/bin/caddy \
 && rm -rf /tmp/caddy

CMD ["caddy", "-quic","-port","8888", "--conf", "/etc/Caddyfile"]