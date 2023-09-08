# 使用官方的 Golang 镜像作为基础镜像
FROM golang:1.20 AS builder
# 使用适用于 ARM 64 位架构的 Golang 镜像
# FROM arm64v8/golang:1.20 AS builder

# 设置工作目录
WORKDIR /app

# 将项目文件复制到容器中
COPY . .

# 编译应用程序，注意替换 main.go 为您的入口文件
RUN go build -o xapi-backend main.go

# 使用 Alpine Linux 作为最终的基础镜像
FROM alpine:latest

# 安装 GLIBC 和其他运行时库
RUN apk --no-cache add ca-certificates libc6-compat

# 设置工作目录
WORKDIR /app

# 复制二进制文件从构建阶段的镜像到最终的镜像
COPY --from=builder /app/xapi-backend .

# 拷贝配置文件到容器中
COPY ./conf /app/conf

# 暴露应用程序所监听的端口
EXPOSE 8090

# 启动应用程序
CMD ["./xapi-backend"]
