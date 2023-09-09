# xapi 项目

在线体验地址：[X API]()

## 项目介绍

基于 React + Golang + Gin + Dubbo-go 的 API接口开放调用平台。

管理员可以接入并发布接口，同时统计分析各接口的调用情况；用户可以注册登录并开通接口调用权限，浏览接口以及在线进行调试，并可以通过 客户端SDK 轻松调用接口。

## 项目架构设计

这个项目的架构是一个典型的分布式系统，包括多个独立的子项目，每个子项目具有特定的功能和职责。以下是每个子项目的功能和关系。

1. 前端项目 [项目仓库: xapi-frontend](https://github.com/xiaoxiongmao5/xapi-frontend)
    * 功能：提供用户界面，用户可以登录、查看接口列表、发起接口调用。
2. 后端项目 [项目仓库: xapi-backend](https://github.com/xiaoxiongmao5/xapi-backend)
    * 功能：提供用户认证、接口管理、账号管理、接口调用功能。
3. 网关项目 [项目仓库: xapi-gateway](https://github.com/xiaoxiongmao5/xapi-gateway)
    * 功能：统一鉴权、限流、路由转发、统一日志、接口染色、统一业务处理等。
4. 客户端SDK项目 [项目仓库: xapi-clientsdk](https://github.com/xiaoxiongmao5/xapi-clientsdk)
    * 功能：封装了对模拟接口项目的调用，提供了简化的API以便其他项目使用。
5. 模拟接口项目 (这里我直接使用的是在线的第三方服务)
    * 功能：模拟第三方接口，供客户端SDK调用，用于开发和测试。

## 接口调用流程图示
![image](https://github.com/xiaoxiongmao5/xapi-backend/assets/25204083/fa3513b8-f58c-4fa7-ac6a-e4917749f636)


## 技术栈

### 后端技术栈
* 主语言：Golang
* 框架：Gin
* 数据库：Mysql8.0、Redis
* 注册中心：Nacos
* RPC远程调用：Dubbo-go
* 微服务网关：Gin
* 接口文档生成：swagger
* 技术设计：API签名认证

### 前端技术栈

* 开发框架：React、Umi 4
* 脚手架：Ant Design Pro
* 组件库：Ant Design、Ant Design Components
* 语法扩展：TypeScript、Less
* 打包工具：Webpack
* 代码规范：ESLint、StyleLint、Prettier
* 图表展示：Echats
* 接口代码生成：OpenAPI

# xapi-backend (API开放平台-后端)

## 项目概述

本项目是一个面向开发者的API平台，提供API接口供开发者调用。用户通过注册登录，可以开通接口调用权限，并可以浏览和调用接口。每次调用都会进行统计，用户可以根据统计数据进行分析和优化。管理员可以发布接口、下线接口、接入接口，并可视化接口的调用情况和数据。

## 项目本地启动

```bash
go mod tidy
go run main.go
```
该项目的接口文档地址：http://<该项目部署站点>/swagger/index.html

## 运行项目中的单元测试

```bash
go test -v ./
go clean -testcache //清除测试缓存
```

## 关于 RPC 远程调用

该项目内的部分业务使用了dubbo-go 框架的rpc远程调用模式。

* 该项目角色是提供方（Provide）

* 配置文件位置：/conf/dubbogo.yaml

* 具体业务为为： `获得用户信息[GetInvokeUser]`、`获得接口信息[GetInterfaceInfoByIdReq]`、`更新接口调用次数[InvokeCount]` 。

### 相关命令

1. 运行注册中心nacos：[见文档](https://blog.csdn.net/trinityleo5/article/details/132622712?spm=1001.2014.3001.5502)

2. 在项目根目录下运行下面命令，生成 rpc 相关go文件，然后共享 ./rpc_api 文件夹 给远程调用方的项目使用
    ```bash
    protoc --go_out=. --go-triple_out=. ./api.proto
    ```

## 其他补充

* 使用 `swag` 生成接口文档命令
    ```bash
    swag fmt
    swag init 
    ```

* 在使用swag生成接口文档后，运行下面指令：将swagger.json挂在服务器上。然后提供地址 `http://<挂载的服务器IP:端口>/swagger.json` 给前端。前端可以在此基础上使用插件自动生成接口请求代码。
比如在[xapi-frontend项目中]运行`npm run openapi`（"openapi": "max openapi"） 可生成对应的接口函数。

    ```bash
    cd ./docs
    python3 -m http.server
    ```

* 使用 `sqlc` 自动生成后端 CRUD 基础代码的命令
    ```bash
    sqlc generate
    ```
