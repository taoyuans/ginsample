package cmd

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"time"

	"ginsample/component/models"
	"ginsample/component/routers"
	cconfig "ginsample/config"
	"ginsample/lib/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func run() {

	err := cconfig.LoadConfig()
	if err != nil {
		panic(err)
	}

	// 初始化mysql
	gormDB := setMySqlDBConfig()
	sqlDB, err := gormDB.DB()
	if err != nil {
		panic(err)
	}

	defer sqlDB.Close()

	// 初始化 Redis
	redis := setRedisDBConfig()

	if initData {
		models.InitData(gormDB)
		os.Exit(1)
	} else {
		if err := setMiddlewareAndRouter(gormDB, redis).Run(":" + cconfig.ConfigValue.HttpPort); err != nil {
			panic(err)
		}
	}
}

// 设置数据库配置
func setMySqlDBConfig() *gorm.DB {
	// 建立数据库连接、
	gormDB, err := gorm.Open(mysql.Open(cconfig.ConfigValue.Mysql.DataSourceName), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		panic(err)
	}

	// 根据需求开启AutoMigrate
	gormDB.AutoMigrate(&models.User{})

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(cconfig.ConfigValue.Mysql.MaxIdleConn)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(cconfig.ConfigValue.Mysql.MaxOpenConn)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Duration(cconfig.ConfigValue.Mysql.MaxConnLifeTime) * time.Second)

	return gormDB
}

// 设置数据库配置
func setRedisDBConfig() *redis.Client {
	// 建立数据库连接、
	rdb := redis.NewClient(&redis.Options{
		Addr:     cconfig.ConfigValue.Redis.DataSourceName,
		Password: cconfig.ConfigValue.Redis.Password,
		DB:       cconfig.ConfigValue.Redis.DB,
	})

	return rdb
}

// 设置中间件以及初始化路由
func setMiddlewareAndRouter(gormDB *gorm.DB, redis *redis.Client) *gin.Engine {
	gin.SetMode(mode)

	r := gin.New()

	r.Static("/static", "static")

	r.Use(gin.Recovery())

	// trace_id的创建
	r.Use(middleware.SetRequestID())

	// 跨域设定
	r.Use(cors.Default())
	// r.Use(middleware.Cors())

	// 添加DB连接
	r.Use(middleware.SetMysqlMiddleware(gormDB))
	r.Use(middleware.SetRedisMiddleware(redis))

	// 添加日志
	r.Use(middleware.SetLogMiddleWare(
		path.Join(cconfig.ConfigValue.Log.LogFilePath, cconfig.ConfigValue.Log.LogFileName),
		cconfig.ConfigValue.Log.LogMaxSize,
		cconfig.ConfigValue.Log.LogMaxBackups,
		cconfig.ConfigValue.Log.LogMaxAge,
		cconfig.ConfigValue.Log.LogLevel))

	// 每次接口调用需不需要做 jwt token 的有效性验证
	if cconfig.ConfigValue.UseJwtAuth {
		r.Use(middleware.CheckToken())
	} else {
		r.Use(middleware.UserClaimMiddelware())
	}

	// 设置接口的Router
	routers.SetBaseRouters(r)
	routers.SetUserRouters(r)

	return r
}

// 获取环境变量信息
func GetEnvDefaultString(key, defaultVal string) string {
	val, ex := os.LookupEnv(key)
	if !ex {
		return defaultVal
	}
	return val
}

func GetEnvDefaultUint64(key string, defaultVal uint64) uint64 {
	val, ex := os.LookupEnv(key)
	if !ex {
		return defaultVal
	}
	num, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		fmt.Printf("GetEnv[%s] err:%v\n", key, err)
		return uint64(0)
	}
	return num
}

func GetEnvDefaultBool(key string, defaultVal bool) bool {
	val, ex := os.LookupEnv(key)
	if !ex {
		return defaultVal
	}
	b, err := strconv.ParseBool(val)
	if err != nil {
		fmt.Printf("GetEnv[%s] err:%v\n", key, err)
		return false
	}
	return b
}

func GetEnvDefaultFloat64(key string, defaultVal float64) float64 {
	val, ex := os.LookupEnv(key)
	if !ex {
		return defaultVal
	}
	b, err := strconv.ParseFloat(val, 64)
	if err != nil {
		fmt.Printf("GetEnv[%s] err:%v\n", key, err)
		return float64(0)
	}
	return b
}
