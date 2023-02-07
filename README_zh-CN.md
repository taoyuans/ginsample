# ginsample

<!-- [![Build Status](https://github.com/wenjianzhang/go-admin/workflows/build/badge.svg)](https://github.com/go-admin-team/go-admin)
[![Release](https://img.shields.io/github/release/go-admin-team/go-admin.svg?style=flat-square)](https://github.com/go-admin-team/go-admin/releases) -->
[![License](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/taoyuans/ginsample)

[English](https://github.com/taoyuans/ginsample/blob/main/README.md) | 简体中文  

- Gin + gorm + mysql(sqlite)

## Directory

```
ginsample 
    ├── component    --> # 业务代码
    │   ├── apis     --> # 接口
    │   │   ├── xxx.go
    │   │   └── xxx.go
    │   ├── models   --> # 详细
    │   │   ├── xxx.go
    │   │   └── xxx.go
    │   └── service  --> # 不推荐使用，简单的用models
    │   │   ├── xxx.go
    │   │   └── xxx.go
    ├── config       --> # 配置文件
    │   ├── dev   
    │   │   └── config.yml
    │   ├── prod    
    │   │   └── config.yml
    │   ├── test    
    │   │   └── config.yml
    │   ├── config.yml   
    │   └── config.go   
    ├── lib          --> # lib模块
    └── main.go
```

## 开始

```
//获取代码
go get github.com/taoyuans/ginsample

//切换目录
cd $GOPATH/src/github.com/taoyuans/ginsample

//gomod 配置
export GO111MODULE=on && export GOPROXY=https://goproxy.cn
go mod tidy

//初始化数据
go run main.go -mode=init -env=dev
//运行服务
go run main.go -mode=api -env=prod
```

访问: <http://127.0.0.1:9001/>

## test

```
go test ./component/apis

go test ./component/models
```

## docker run

#### 普通方式

```
docker build -f dockerfile -t ginsample-image:latest .
docker run -d --name ginsample-container -i -e APP_ENV=prod -p 40001:9001 ginsample-image:latest
```

#### 使用alpine(需要对项目和配置有一定的了解)

```
docker build -f dockerfile_alpine -t ginsample-image:latest .
docker run -d --name ginsample-container --network mingxie-network -v /Users/limingxie/volumes/ginsample/logs:/go/bin/logs:rw -i -e APP_ENV=prod -p 40001:9001 ginsample-image:latest
```

## Import

- web framework: [gin framework](https://github.com/gin-gonic/gin)
- orm tool: [gorm](https://gorm.io/)
- logger : [logrus](https://github.com/sirupsen/logrus)
- configuration tool: [viper](https://github.com/spf13/viper)
<!-- - validator: [govalidator](github.com/asaskevich/govalidator)
- utils: <https://github.com/pangpanglabs/goutils> -->

## References

<https://github.com/pangpanglabs/echosample>  
<https://github.com/go-admin-team/go-admin>

## 🔑 License

[MIT](https://github.com/taoyuans/ginsample/blob/main/LICENSE)

Copyright (c) 2023 li_mingxie

----------------------------------------------

欢迎大家的意见和PR  

`email: li_mingxie@163.com`
