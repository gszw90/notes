# 概述

是一个构建镜像的文本文件，是由一条条构建镜像的命令与参数构成的脚本

## 基础概念

1. 命令都是大写字母，且至少有一个参数
2. 命令从上到下依次执行
3. 使用#进行注释
4. 每条指令都会创建一个新的镜像层并对镜像进行提交

# 关键字

| 命令         | 说明                     | 其他                                                                                                                                  |
|------------|------------------------|-------------------------------------------------------------------------------------------------------------------------------------|
| FROME      | 镜像的基础父类，再文件的第一行        |                                                                                                                                     |
| MAINTAINER | 维护者信息，一般包括姓名与邮箱        |                                                                                                                                     |
| RUN        | 容器构建时执行的命令             | 有两种格式，一种是shell，一种是exec                                                                                                              |
| EXPOSE     | 对外暴露的端口                |                                                                                                                                     |
| WORKDIR    | 指定容器创建完成后的工作目录         |                                                                                                                                     |
| USER       | 指定运行容器的用户，如果不指定，默认root |                                                                                                                                     |
| ENV        | 设置容器的环境变量              |                                                                                                                                     |
| VOLUME     | 容器卷，宿主机与容器的文件映射        |                                                                                                                                     |
| ADD        | 将宿主机的文件拷贝到容器内          | 能自动解压压缩文件，也能把远程的url地址文件下载下来拷贝到容器内                                                                                                   |
| COPY       | 将宿主机的文件拷贝到容器内          | 只是拷贝操作，不会解压或者下载远程文件                                                                                                                 |
| CMD        | 指定容器启动时需要执行的命令         | 支持shell与exec格式的命令，<br/>如果指定了ENDPOINT,CMD里面的参数会作为endpoint的参数来使用<br/>可以有多个CMD，但是只有最后一个会生效<br/>如果docker run时后面有自定义的参数命令，该参数命令会覆盖掉CMD命令 |
| ENDPOINT   | 指定容器启动时需要执行的命令         | 不会被docker run 后面的命令覆盖，docker run后面的值会被当做是endpoint的参数来使用                                                                             |
# 构建镜像
在Dockerfile同级目录下执行
```bash
dcoker build -t image_name:tag .
```
## 例子
1. 准备go文件
```go
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync/atomic"
)

var counter int64

func main() {
	port := os.Getenv("APP_PORT")
	if len(port) == 0 {
		port = "8089"
	}

	http.HandleFunc("/count", func(writer http.ResponseWriter, request *http.Request) {
		val := atomic.AddInt64(&counter, 1)
		writer.Write([]byte(fmt.Sprintf("counter:%v", val)))
	})
	log.Printf("start app at port:%v\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
	if err != nil {
		panic(fmt.Sprintf("fail to run http server:%s", err))
	}
}
```
2. 准备Dockerfile文件
```Dockerfile
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
```

3. 镜像制作
```bash
 docker build -t go_app:v1.0.0 .
```