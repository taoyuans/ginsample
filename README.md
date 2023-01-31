# ginsample

[![License](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/taoyuans/ginsample)

English | [简体中文](https://github.com/taoyuans/ginsample/blob/main/README_zh-CN.md)

Gin + gorm + mysql(sqlite)

## Directory

```
ginsample 
    ├── component    --> # business code
    │   ├── apis     --> # interface
    │   │   ├── xxx.go
    │   │   └── xxx.go
    │   ├── models   --> # source code
    │   │   ├── xxx.go
    │   │   └── xxx.go
    │   └── service  --> # not Recommended, use models
    │   │   ├── xxx.go
    │   │   └── xxx.go
    ├── config       --> # config file
    │   ├── dev   
    │   │   └── config.yml
    │   ├── prod    
    │   │   └── config.yml
    │   ├── test    
    │   │   └── config.yml
    │   ├── config.yml   
    │   └── config.go   
    ├── lib          --> # lib module
    └── main.go
```

## Getting Started

```
//get source
go get github.com/taoyuans/ginsample

//dir
cd $GOPATH/src/github.com/taoyuans/ginsample

//gomod option
export GO111MODULE=on && export GOPROXY=https://goproxy.cn
go mod tidy

//init data
go run main.go -mode=init -env=dev
//run
go run main.go -mode=api -env=prod
```

Visit <http://127.0.0.1:9001/>

## test

```
go test ./component/apis

go test ./component/models
```

## docker run

```
docker build -f dockerfile -t ginsample-image:latest .
docker run -d --name ginsample-container -i -e APP_ENV=prod -p 40001:9001 ginsample-image:latest
```

## Import

- web framework: [gin framework](https://github.com/gin-gonic/gin)
- orm tool: [gorm](https://gorm.io/)
- logger : [logrus](https://github.com/sirupsen/logrus)
- configuration tool: [viper](https://github.com/spf13/viper)

## References

<https://github.com/pangpanglabs/echosample>  
<https://github.com/go-admin-team/go-admin>

## 🔑 License

[MIT](https://github.com/taoyuans/ginsample/blob/main/LICENSE)

Copyright (c) 2023 li_mingxie

----------------------------------------------

Welcome comments and PR  

`email: li_mingxie@163.com`
