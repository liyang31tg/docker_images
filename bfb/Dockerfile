FROM golang
RUN go get github.com/gin-gonic/gin \
    && go get github.com/go-redis/redis \
	&& go get github.com/gin-gonic/contrib/sessions \
	&& go get github.com/go-sql-driver/mysql \
	&& cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime  #修改服务器上的时区