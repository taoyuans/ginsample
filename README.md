## Directory

```
ginsample-api 
    ├── cmd          --> # app start
    ├── component    --> # business module
    │   ├── apis     --> # interface
    │   │   ├── xxx.go
    │   │   └── xxx.go
    │   └── common  --> # common code about business
    │   │   ├── xxx.go
    │   │   └── xxx.go
    │   ├── models   --> # source code
    │   │   ├── xxx.go
    │   │   └── xxx.go
    │   ├── routers   --> # router 
    │   │   ├── xxx.go
    │   │   └── xxx.go
    │   └── service  --> # business code
    │   │   ├── xxx.go
    │   │   └── xxx.go
    ├── config       --> # load nacos config 
    ├── lib          --> # lib module
    └── main.go
```

## Getting Started

```
//gomod option
export GO111MODULE=on && export GOPROXY=https://goproxy.cn
go mod tidy

//init data
go run main.go -INIT=true 
//run
go run main.go
```

Visit <http://127.0.0.1:33888/>

## test

```
go test ./component/apis

go test ./component/models
```

## docker run

#### normal

```
docker build -f dockerfile_normal -t ginsample-api-image:latest .

docker run -d --name ginsample-api-container -i -e MODE=release -p 33888:33888 ginsample-api-image:latest
```

#### use alpine

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

----------------------------------------------

`email: li_mingxie@163.com`
