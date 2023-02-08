package middleware

import (
	"context"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

func SetLogMiddleWare(logFilePath, logFileName string) gin.HandlerFunc {
	// create log filename
	fileName := path.Join(logFilePath, logFileName)
	// open file
	file, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	// file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open/write log file failed.", err)
		return nil
	}
	// 实例化
	logger := logrus.New()
	// set log level
	logger.SetLevel(logrus.DebugLevel)
	// set out
	logger.Out = file
	// set rotatelogs for file segmentation
	logWriter, err := rotatelogs.New(
		// 分割后的文件名称
		fileName+".%Y%m%d.log",
		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(fileName),
		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(7*24*time.Hour), //以hour为单位的整数
		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(1*time.Hour),
	)
	// set hook
	writerMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	//set logrus hook
	logger.AddHook(lfshook.NewHook(writerMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}))

	return func(c *gin.Context) {
		req := c.Request
		c.Request = req.WithContext(context.WithValue(req.Context(), "Logger", logger))
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		// 打印日志
		logger.WithFields(logrus.Fields{
			"status_code": c.Writer.Status(),
			"latencyTime": endTime.Sub(startTime),
			"client_ip":   c.ClientIP(),
			"req_method":  c.Request.Method,
			"req_uri":     c.Request.RequestURI,
		}).Info()

	}
}
