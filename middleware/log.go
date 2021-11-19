package middleware

import (
	"fmt"
	"time"

	"github.com/CodeHanHan/ferry-backend/pkg/logger"
	"github.com/gin-gonic/gin"
)

const (
	Authorization = "Authorization"
	RequestIdKey  = "X-Request-Id"
	xForwardedFor = "X-Forwarded-For"
)

func Logger() gin.HandlerFunc {
	l := logger.NewLogger()
	l.HideCallstack()
	return func(c *gin.Context) {
		raw := c.Request.URL.RawQuery

		t := time.Now()

		// process request
		c.Next()

		latency := time.Since(t)
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		path := c.Request.URL.Path

		if raw != "" {
			path = path + "?" + raw
		}

		logStr := fmt.Sprintf(" %3d | %v | %s | %s %s %s ",
			statusCode,
			latency,
			clientIP, method,
			path,
			c.Errors.String(),
		)

		switch {
		case statusCode >= 400 && statusCode <= 499:
			l.Warn(c, logStr)
		case statusCode >= 500:
			l.Error(c, logStr)
		default:
			l.Info(c, logStr)
		}
	}
}
