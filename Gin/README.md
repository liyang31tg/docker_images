#Gin的docker镜像
> `Dockerfile` ,只要将源码挂载到/app目录下面，并且入口文件从main.go就可以了
```
FROM registry.cn-hangzhou.aliyuncs.com/lizi/gin
WORKDIR /app
CMD ["go","run","main.go"]
```
