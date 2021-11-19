package app

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Error(c *gin.Context, code int, format string, values ...interface{}) {
	errMsg := http.StatusText(code)

	var msg string
	if len(values) > 0 {
		msg = fmt.Sprintf("%s | "+format, errMsg, values)
	} else {
		msg = errMsg + ": " + format
	}
	c.String(code, msg)
}

func OK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}
