#Redis的docker镜像
```
redis:
    image: registry.cn-hangzhou.aliyuncs.com/lizi/redis
    ports:
      - "6378:6379"
    volumes:
      - /Users/user/Desktop/redisdata:/data 	   //持久化文件的挂载
      - ./config/redis.conf:/etc/redis/redis.conf //配置文件的挂载
```