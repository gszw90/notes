# DockerCompose

## 安装

### ubuntu

```bash
# 1.安装依赖
sudo apt-get install apt-transport-https ca-certificates curl software-properties-common
# 2. 添加秘钥
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
# 3. 添加仓库
sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
# 4. 更新索引
sudo apt-get update
# 5. 安装
sudo apt-get install docker-compose
```

### 二进制

进入[github下载页面](https://github.com/docker/compose/releases)选择对应版本，下载下来，添加执行权限

```bash
sudo curl -L "https://github.com/docker/compose/releases/download/v2.27.1/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
```

## 常用命令

执行命令需要当前目录下有docker-compose.yml文件，或者需要指定配置文件

```bash
# 创建并运行容器
# -f:指定配置文件，默认docker-compose.yml
# -d:后台运行容器
docker-compose up
docker-compose -f docker-compose.dev.yaml up -d
# 重构容器
# --force-rm：删除构建过程中的临时容器
# --pull：尝试拉取最先版本镜像
docker-compose build
# 启动已经存在的容器
docker-compose start
# 重启
docker-compose restart
# 停止容器
dcoker-compose stop
# 停止并删除docker-compose创建的容器，网络，数据卷等
dcoekr-compose down
# 查看管理的容器状态
docker-compose ps
# 查看容器日志
docker-compose logs
# 删除已停止的容器
docker-compose rm
# 查看容器进程
docker-compose top
# 查看容器资源使用情况
docker-compose stats
# 验证配置文件是否合法
docker-compose config
```

docker-compose.yml

```yml
# compose版本
version: "3"
# 服务列表
services:
  # 自定义服务名称，如果没有配置container_name，容器名就是这个
  my_ubuntu:
    # 基于Dockerfile构建服务时需要配置的参数
    build:
      # 指定Dockerfile文件的路径，.表示当前目录
      context: .
      # 指定Dockerfile文件的名字，默认Dockerfile
      dockerfile: Dockerfile
    # 基础镜像，如果使用了build，就不需要设置image
    image: ubuntu:latest
    # 容器名称，不配置就使用上面的服务名称
    container_name: my_ubuntu
    # 端口映射，用于映射宿主机与容器间的端口
    ports:
      # 宿主机的8089端口与容器的8089端口进行映射
      - 8089:8089
      # 宿主机的一个随机端口与容器的3000端口进行映射
      - 3000
    # 端口暴露，但是不会把端口映射到宿主机，用于容器间的端口通信
    expose:
      - 8089
    # 目录挂载，既可以挂载宿主机的文件或者目录，也可以挂载数据卷
    volumes:
      - ./data:/data
    environment:
      APP_PORT: 8089
      APP_NAME: my_app
    # 指定依赖，确保当前服务在依赖服务启动后才能启动
    depends_on:
      - db
    # 指定网络
    networks:
      - app_network
    # 容器启动后执行的命令
    command: "echo success"
  db:
    image: mysql:8.4.0
    container_name: db
    ports:
      - 13306:3306
    # 挂载数据卷
    volumes:
      - db_data:/var/lib/mysql/data
    networks:
      - app_network

# 自定义网络配置
networks:
  # 网络名称
  app_network:
    # 网络类型：bridge，host，none
    driver: bridge
# 自定义数据卷
volumes:
  db_data:
```