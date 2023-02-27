# kilikili
## config.ini
```
# debug开发模式,release生产模式
[service]
AppMode = debug
HttpPort = :3000
# 运行端口号 3000端口

[mysql]
Db = mysql
DbHost = 127.0.0.1
# mysql ip地址
DbPort = 3306
# mysql 端口号
DbUser = root
# mysql 用户名
DbPassWord = 123456
# mysql 密码
DbName = kilikili
# mysql 名字

[redis]
RedisDb = redis
RedisAddr = localhost:6379
# redis ip地址和端口号
RedisPw =
# redis 密码
RedisDbName = 1
# redis 名字
```
