>`docker-compose.yml`
```
caddy:
  image: registry.cn-hangzhou.aliyuncs.com/lizi/caddy
  ports:
    - "80:80"
  volumes:
    - /Users/liyang/Downloads/caddy_ubuntu/Caddyfile:/etc/Caddyfile //替换
  links:
    - webapp1
    - webapp2
```

>`Caddyfile`
```
local.boqiao919.com:80 {
	index /caddy/index.html
	tls off
	proxy / webapp1:8080 {
		policy ip_hash 
		transparent
	}
	proxy / webapp2:8080 {
		policy ip_hash 
		transparent
	}
}
```
