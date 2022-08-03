# Npool ledger manager

[![Test](https://github.com/NpoolPlatform/service-template/actions/workflows/main.yml/badge.svg?branch=master)](https://github.com/NpoolPlatform/service-template/actions/workflows/main.yml)

[目录](#目录)
- [功能](#功能)
- [命令](#命令)
- [步骤](#步骤)
- [最佳实践](#最佳实践)
- [关于mysql](#关于mysql)

-----------
### 功能
### 功能
- [x] 创建service template
- [x] 封装日志库
- [x] 统一service cli框架
- [x] 集成cli框架(https://github.com/urfave/cli)
- [x] 集成http server框架(https://github.com/go-chi/chi.git 不需要封装)
- [x] 集成http client框架(https://github.com/go-resty/resty 不需要封装)
- [x] 集成consul注册与发现
- [x] 全局主机环境参数解析
- [x] 集成apollo配置中心(https://github.com/philchia/agollo.git)
- [x] 集成redis访问
- [x] 集成mysql访问框架(https://github.com/ent/ent)
* [x] 集成版本信息
* [x] 集成rabbitmq访问
* [x] 完善rabbitmq API
* [x] 生成docker镜像
* [x] 发布docker镜像
* [x] 将服务部署到k8s集群
* [x] 将服务api通过traefik-internet ingress代理，供外部应用调用(视服务功能决定是否需要)
* [ ] ingress中服务相关api的traefik规则定义
* [x] 集成GRPC
* [x] 添加GRPC proto编译支持

### 命令
* make init ```初始化仓库，创建go.mod```
* make verify ```验证开发环境与构建环境，检查code conduct```
* make verify-build ```编译目标```
* make test ```单元测试```
* make generate-docker-images ```生成docker镜像```
* make service-template ```单独编译服务```
* make service-template-image ```单独生成服务镜像```
* make deploy-to-k8s-cluster ```部署到k8s集群```

### 步骤
* 在github上将模板仓库https://github.com/NpoolPlatform/service-template.git import为https://github.com/NpoolPlatform/my-service-name.git
* git clone https://github.com/NpoolPlatform/my-service-name.git
* cd my-service-name
* mv cmd/service-template cmd/my-service
* 修改cmd/my-service/main.go中的serviceName为My Service
* mv cmd/my-service/ServiceTemplate.viper.yaml cmd/my-service/MyService.viper.yaml
* 将cmd/my-service/MyService.viper.yaml中的内容修改为当前服务对应内容
* 修改Dockerfile和k8s部署文档为当前服务对应内容
  * grep -rb "service template" ./*
  * grep -rb "ServiceTemplate" ./*
  * grep -rb "Service Template" ./*
  * grep -rb "service_template" ./*
  * grep -rb "service-template" ./*
  * grep -rb "servicetmpl" ./*
  * 修改cmd/my-service/k8s中的三个yaml文件，包含端口，服务名字
  * grep -rb "serviceid"，并使用uuid生成新值替换

### 最佳实践
* 每个服务只提供单一可执行文件，有利于docker镜像打包与k8s部署管理
* 每个服务提供http调试接口，通过curl获取调试信息
* 集群内服务间direct call调用通过服务发现获取目标地址进行调用
* 集群内服务间event call调用通过rabbitmq解耦

### 关于mysql
* 创建app后，从app.Mysql()获取本地mysql client
* [文档参考](https://entgo.io/docs/sql-integration)

### API规范(假设目标对象为Object)
# 前端接口或App管理员接口
* 创建 ```CreateObject```
* 批量创建 ```CreateObjects```
* 更新 ```UpdateObject```
* 更新指定域 ```UpdateObjectFields```
* 原子加指定域 ```AddObjectFields```
* 用ID查询 ```GetObject```
* 条件查询单条 ```GetObjectOnly```
* 条件查询多条 ```GetObjects```
* 条件计数 ```CountObjects```
* 查询ID是否存在 ```ExistObject```
* 查询条件是否存在 ```ExistObjectConds```
* ID删除 ```DeleteObject```

# 大后台管理员接口
* 为其他App创建 ```CreateAppObject```
* 批量为其他App创建 ```CreateAppObjects```
* 查询其他App多条 ```GetAppObjects```

# 注意事项
* 网关对请求的http header字段做如下转换(request为http请求body)
  * X-App-ID -> request.AppID,
  * X-App-ID -> request.Info.AppID, 如果request携带Info字段
  * X-User-ID -> request.UserID,
  * X-User-ID -> request.Info.UserID, 如果request携带Info字段
  * X-Lang-ID -> request.LangID,
  * X-Lang-ID -> request.Info.LangID, 如果request携带Info字段
  * X-App-User-Token -> request.Token,
  * X-App-User-Token -> request.Info.Token, 如果request携带Info字段

* 对于通过条件查询的前端API或App管理员API，其request中应总是包含AppID字段，实现时应该总是将该AppID字段添加到查询条件
* 对于批量创建的前端或App管理员API，其request中应总是包含AppID字段，实现时应该总是将该AppID字段覆盖request.Info或request.Infos中的AppID
* 更新API的实现不应更新AppID域
* 对于通过条件查询或创建其他App数据的大后台管理员API，其request中应总是包含TargetAppID，实现时应该总是将TargetAppID字段覆盖request.Info或request.Infos中的AppID
