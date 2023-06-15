## Directory

```
ginsample-api 
    ├── cmd          --> # 启动程序
    ├── component    --> # 业务代码
    │   ├── apis     --> # 接口定义
    │   │   ├── xxx.go
    │   │   └── xxx.go
    │   └── common  --> # 业务相关的公用代码
    │   │   ├── xxx.go
    │   │   └── xxx.go
    │   ├── models   --> # struct和数据库连接信息
    │   │   ├── xxx.go
    │   │   └── xxx.go
    │   ├── routers   --> # 路由
    │   │   ├── xxx.go
    │   │   └── xxx.go
    │   └── service  --> # 具体业务代码
    │   │   ├── xxx.go
    │   │   └── xxx.go
    ├── config       --> # 读取nacos配置
    ├── lib          --> # lib模块, 业务无关的代码
    └── main.go
```

## 开始

```
//gomod 配置
export GO111MODULE=on && export GOPROXY=https://goproxy.cn
go mod tidy

//初始化数据
go run main.go -INIT=true 
//运行服务
go run main.go
```

访问: <http://127.0.0.1:33888/>

## test

```
go test ./component/apis

go test ./component/models
```

## docker run

#### 普通方式

```
docker build -f dockerfile_normal -t ginsample-api-image:latest .

docker run -d --name ginsample-api-container -i -e MODE=release -p 33888:33888 ginsample-api-image:latest
```

#### 使用alpine

```
docker build -f dockerfile -t ginsample-api-image:latest .

docker run -d --name ginsample-api-container \
-v /Users/limingxie/volumes/logs/ginsample-api:/app/logs:rw \
-i -e MODE=release -p 33888:33888 ginsample-api-image:latest
```

## Import

- web framework: [gin framework](https://github.com/gin-gonic/gin)
- orm tool: [gorm](https://gorm.io/)
- logger : [zerolog](https://github.com/rs/zerolog)
- configuration tool: [nacos](https://github.com/nacos-group/nacos-sdk-go)

----------------------------------------------

`email: li_mingxie@163.com`
