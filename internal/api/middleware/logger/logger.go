package logger

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func New() gin.HandlerFunc {

	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		clientIP := c.ClientIP()
		statusCode := c.Writer.Status()
		if clientIP == "" {
			clientIP = c.Request.Header.Get("Origin")
		}
		switch statusCode {
		case 500, 501, 503:
			go logrus.Errorf("| %3d | %13v | %15s | %s | %s |", c.Writer.Status(), latencyTime, clientIP, c.Request.Method, c.Request.RequestURI)
		case 400, 404, 403:
			go logrus.Errorf("| %3d | %13v | %15s | %s | %s |", c.Writer.Status(), latencyTime, clientIP, c.Request.Method, c.Request.RequestURI)
		}
		logrus.SetOutput(os.Stdout)
	}
}
