package app

import (
	"fmt"
	"net/http"

	"github.com/CodeHanHan/ferry-backend/pkg/validator"
	"github.com/gin-gonic/gin"
)

func InternalServerError(c *gin.Context) {
	c.String(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
}

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

func ErrorParams(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, validator.Translate(err))
}

func OK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}
