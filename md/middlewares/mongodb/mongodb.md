# mongodb

## 单例

### 准备docker-compose.yaml
```yaml
version: '3.8'

services:
  mongo6:
    image: mongodb/mongodb-community-server:6.0.15-ubi8
    container_name: mongo6
    environment:
      - TZ=Asia/Shanghai
      - MONGODB_INITDB_ROOT_USERNAME=zeng
      - MONGODB_INITDB_ROOT_PASSWORD=zeng
    ports:
      - "27017:27017"
    volumes:
      - ./data/db:/data/db
    networks:
      - mongo_network

networks:
  mongo_network:
    driver: bridge
```
### 启动容器
```bash
docker-compose up -d
```


## 后面的失败了

## 准备工作
### 生成秘钥
```bash
openssl rand -base64 756 > ./mongo-keyfile
```
### 配置docker-compose

```dockerfile
FROM mongo:6.0.15

COPY ./mongo-keyfile /opt/keyfile/mongo-keyfile
RUN  chmod 600 /opt/keyfile/mongo-keyfile
```

```yaml
version: "3"

services:
  mongodb6-1:
    #    image: "mongo:6.0.15"
    build:
      context: .
      dockerfile: Dockerfile
    container_name: mongodb6-1
    ports:
      - "27017:27017"
    environment:
      MONGO_INITDB_ROOT_USERNAME: moment
      MONGO_INITDB_ROOT_PASSWORD: moment
    volumes:
      - ./mongo-keyfile:/opt/keyfile/mongo-keyfile:ro
      - ./data/db1:/data/db
    command: mongod --replSet rs0 --auth --keyFile /opt/keyfile/mongo-keyfile
    networks:
      - mongo-network

  mongodb6-2:
    image: "mongo:6.0.15"
#    build:
#      context: .
#      dockerfile: Dockerfile
    container_name: mongodb6-2
    depends_on:
      - mongodb6-1
    environment:
      MONGO_INITDB_ROOT_USERNAME: moment
      MONGO_INITDB_ROOT_PASSWORD: moment
    volumes:
      - ./mongo-keyfile:/opt/keyfile/mongo-keyfile:ro
      - ./data/db2:/data/db
    command: mongod --replSet rs0 --auth --keyFile /opt/keyfile/mongo-keyfile
    networks:
      - mongo-network

  mongodb6-3:
    image: "mongo:6.0.15"
#    build:
#      context: .
#      dockerfile: Dockerfile
    container_name: mongodb6-3
    depends_on:
      - mongodb6-1
    environment:
      MONGO_INITDB_ROOT_USERNAME: moment
      MONGO_INITDB_ROOT_PASSWORD: moment
    volumes:
      - ./mongo-keyfile:/opt/keyfile/mongo-keyfile:ro
      - ./data/db3:/data/db

    command: mongod --replSet rs0 --auth --keyFile /opt/keyfile/mongo-keyfile --oplogSize 128
    networks:
      - mongo-network

networks:
  mongo-network:
    driver: bridge
```

## 启动集群
1. 启动容器
```bash
docker-compose up -d
```
2. 进入容器
```bash
docker exec -it  mongodb6-1 bash
```
3. 在容器内执行
```bash
mongosh "mongodb://moment:moment@localhost:27017/?authSource=admin" 
```