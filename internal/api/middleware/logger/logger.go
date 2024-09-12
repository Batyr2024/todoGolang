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

		go func() {
			startTime := time.Now()
			c.Next()
			endTime := time.Now()
			latencyTime := endTime.Sub(startTime)
			clientIP := c.ClientIP()
			if clientIP == "" {
				clientIP = c.Request.Header.Get("Origin")
			}
			switch c.Writer.Status() {
			case 500, 501, 503:
				logrus.Fatalf("| %3d | %13v | %15s | %s | %s |", c.Writer.Status(), latencyTime, clientIP, c.Request.Method, c.Request.RequestURI)
			case 400, 404, 403:
				logrus.Errorf("| %3d | %13v | %15s | %s | %s |", c.Writer.Status(), latencyTime, clientIP, c.Request.Method, c.Request.RequestURI)
			}
			logrus.SetOutput(os.Stdout)
		}()
	}
}
