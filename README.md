# CMall 电子商城

#### 此项目已经部署至[CMall](http://cmall.congz.top/#/)，用 golang 实现接口函数，如需要看前端请前往[cmall-vue](https://github.com/congz666/cmall-vue)

## 前言

本人准备大三，拿这个项目学习一些新的技能，如果有错误或者实现不好的地方欢迎 issues

后续会实现支付功能，QQ 第三方登录，重写搜索系统，优化代码和数据库，因为即将开学，后续更新的进度可能会放慢。

如果觉得这个项目不错，您可以右上角 Star 支持一下！谢谢您的支持，您的支持是我完善的动力！

## 项目依赖

本项目采用了一系列 golang 中比较流行的组件来进行开发

- Gin
- Gorm
- mysql
- redis
- godotenv
- jwt-go
- 阿里云 OSS
- 极验 SDK

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
│   ├── sdk
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

- pkg / sdk: 极验 sdk 核心函数

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
export GOPROXY=http://mirrors.aliyun.com/goproxy/
go mod tidy
```

## 运行

```
go run main.go
```

项目运行后启动在 3000 端口
