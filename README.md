# xapi-backend (API开放平台-后端)

## 项目概述

本项目是一个面向开发者的API平台，提供API接口供开发者调用。用户通过注册登录，可以开通接口调用权限，并可以浏览和调用接口。每次调用都会进行统计，用户可以根据统计数据进行分析和优化。管理员可以发布接口、下线接口、接入接口，并可视化接口的调用情况和数据。

## 项目架构


## 项目本地启动

```bash
go mod tidy
go run main.go
```

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

* 在使用swag生成接口文档后，运行下面指令：将swagger.json挂在服务器上，以供前端能运行`npm run openapi`（"openapi": "max openapi"） 生成对应的接口函数

    ```bash
    cd ./docs
    python3 -m http.server
    ```