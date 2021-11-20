package app

import (
	"fmt"
	"net/http"

	"github.com/CodeHanHan/ferry-backend/pkg/logger"
	"github.com/gin-gonic/gin"
)

func Error(c *gin.Context, err error, code int, format string, values ...interface{}) {
	errMsg := http.StatusText(code)

	logger.Error(c, err.Error())
	var msg string
	if len(values) > 0 {
		msg = fmt.Sprintf("%s | "+format, errMsg, values)
	} else {
		msg = errMsg + ": " + format
	}
	c.String(code, msg)
}

func ErrorParams(c *gin.Context, err error) {
	logger.Error(c, "参数验证失败: %v", err.Error())

	c.String(http.StatusBadRequest, err.Error())
}

func OK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}
