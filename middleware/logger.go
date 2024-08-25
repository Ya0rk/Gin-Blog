package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"math"
	"os"
	"time"
)

/*
自定义log日志中间件
https://pkg.go.dev/github.com/sirupsen/logrus@v1.6.0#section-readme

日志分割显示：
file-totatelogs : https://pkg.go.dev/github.com/lestrrat-go/file-rotatelogs#section-readme
lfshook : https://pkg.go.dev/github.com/rifflock/lfshook
*/
func Logger() gin.HandlerFunc {
	// 保存 log 日志的文件路径
	filePath := "log/log"
	linkName := "latest_log.log"

	src, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("open file err:", err)
	}

	logger := logrus.New()

	// 将日志写进文件中
	logger.Out = src

	// 配置日志级别
	logger.SetLevel(logrus.DebugLevel)
	logWriter, _ := rotatelogs.New(
		filePath+"%Y%m%d.log",
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
		rotatelogs.WithLinkName(linkName),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	Hook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logger.AddHook(Hook)

	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		stopTime := time.Since(startTime)
		spendTime := fmt.Sprintf("%d ms", int(math.Ceil(float64(stopTime.Nanoseconds())/float64(1000000))))
		hostName, err := os.Hostname()
		if err != nil {
			hostName = "unknown"
		}

		statusCode := c.Writer.Status()
		clientIp := c.ClientIP()
		userAgent := c.Request.UserAgent()
		dataSize := c.Writer.Size()
		if dataSize < 0 {
			dataSize = 0
		}
		method := c.Request.Method
		path := c.Request.RequestURI

		entry := logger.WithFields(logrus.Fields{
			"HostName":   hostName,
			"statusCode": statusCode,
			"spendTime":  spendTime,
			"clientIp":   clientIp,
			"Method":     method,
			"path":       path,
			"dataSize":   dataSize,
			"Agent":      userAgent,
		})

		// 记录系统内部错误
		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if statusCode >= 500 {
			entry.Error()
		} else if statusCode >= 400 {
			entry.Warn()
		} else {
			entry.Info()
		}
	}
}
