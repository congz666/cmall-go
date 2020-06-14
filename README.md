# CMall（小米商城）

#### 此项目用golang实现接口函数，如需要前端请移步[cmall-vue](https://github.com/congz666/cmall-vue) 

#### 如果有错误或者实现不好的地方欢迎issues

#### 如果觉得这个项目不错，您可以右上角Star支持一下，谢谢

## 项目依赖

- Gin
- Gorm
- mysql
- redis（还没上）
- godotenv



## 目录结构

```
mall-go/
├── api
├── cache
├── conf
├── middleware
├── model
├── serializer
├── server
├── service
└── util
```

- api：用于定义接口函数
- cache：redis相关操作（还没上）
- conf：用于存储配置文件
- middleware：应用中间件
- model：应用数据库模型
- serializer：将数据序列化为json的函数
- server 路由逻辑处理
- service：接口函数的实现
- util：工具函数

## Godotenv

项目在启动的时候依赖以下环境变量，但是在也可以在项目根目录创建.env文件设置环境变量便于使用(建议开发环境使用)

```
MYSQL_DSN="db_user:db_password@/db_name?charset=utf8&parseTime=True&loc=Local" # Mysql连接地址
REDIS_ADDR="127.0.0.1:6379" # Redis端口和地址
REDIS_PW="" # Redis连接密码
REDIS_DB="" # Redis库从0到10
SESSION_SECRET="setOnProducation" # Seesion密钥，必须设置而且不要泄露
GIN_MODE="debug"
```

## Go Mod

本项目使用[Go Mod](https://github.com/golang/go/wiki/Modules)管理依赖。

```
go mod init mall-go
export GOPROXY=http://mirrors.aliyun.com/goproxy/
go run main.go
```

## 运行

```
go run main.go
```

项目运行后启动在3000端口