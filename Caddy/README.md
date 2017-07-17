```
caddy:
  image: registry.cn-hangzhou.aliyuncs.com/lizi/caddy
  ports:
    - "80:80"
  volumes:
    - /Users/liyang/Downloads/caddy_ubuntu/Caddyfile:/etc/Caddyfile
  links:
    - webapp1
    - webapp2
```