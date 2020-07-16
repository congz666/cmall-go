# CMall（小米商城）

#### 此项目用 golang 实现接口函数，如需要看前端或者项目图片请移步[cmall-vue](https://github.com/congz666/cmall-vue)

#### 静态图片存放在阿里云 OSS 上，mysql 存放图片地址

#### 如果有错误或者实现不好的地方欢迎 issues

#### 如果觉得这个项目不错，您可以右上角 Star 支持一下，谢谢

## 说明

本项目采用了一系列 golang 中比较流行的组件来进行开发

登录已由原来的 cookie 和 session 保存状态改为使用 token 验证

静态图片存放在阿里云 OSS 上，mysql 存放图片地址

项目还在完善中... 后续会部署到服务器上

## 项目依赖

- Gin
- Gorm
- mysql
- redis
- godotenv
- jwt-go
- 阿里云 OSS

## 目录结构

```
mall-go/
├── api
├── cache
├── conf
├── middleware
├── model
├── pkg
│	├── e
│	├── util
├── serializer
├── server
└── service

```

- api：用于定义接口函数

- cache：redis 相关操作

- conf：用于存储配置文件

- middleware：应用中间件

- model：应用数据库模型

- pkg / e：封装错误码

- pkg / util：工具函数

- serializer：将数据序列化为 json 的函数

- server 路由逻辑处理

- service：接口函数的实现

## Godotenv

项目在启动的时候依赖以下环境变量，但是在也可以在项目根目录创建.env 文件设置环境变量便于使用(建议开发环境使用)

```
MYSQL_DSN="db_user:db_password@/db_name?charset=utf8&parseTime=True&loc=Local" # Mysql连接地址
REDIS_ADDR="127.0.0.1:6379" # Redis端口和地址
REDIS_PW="" # Redis连接密码
REDIS_DB="" # Redis库从0到10
SESSION_SECRET="setOnProducation" # Seesion密钥，必须设置而且不要泄露
GIN_MODE="debug"
OSS_END_POINT="oss-cn-hongkong.aliyuncs.com"#阿里云oss的配置
OSS_ACCESS_KEY_ID="xxx"
OSS_ACCESS_KEY_SECRET="qqqq"
OSS_BUCKET="lalalal"
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

项目运行后启动在 3000 端口
