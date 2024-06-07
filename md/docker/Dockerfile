FROM golang:1.21.11-alpine3.20

LABEL maintainer="zeng<zeng@163.com>"

ENV APP_PORT 8089
ENV GOPROXY https://goproxy.cn
# 将源码拷贝到容器内
COPY ./src /app
# 设置工作目录
WORKDIR /app
# 打包文件
RUN cd /app && go build -o app .
# 暴露端口
EXPOSE 8089
# 运行程序
CMD "./app"