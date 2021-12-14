package app

import (
	"fmt"
	"net/http"

	"github.com/CodeHanHan/ferry-backend/pkg/validator"
	"github.com/gin-gonic/gin"
)

func InternalServerError(c *gin.Context) {
	Error(c, Err_Internal, nil)
}

func Error(c *gin.Context, code ErrCode, detail interface{}) {
	resp := ErrorResponse{
		Code:    code,
		Message: code.String(),
		Detail:  detail,
	}

	c.JSON(int(code), resp)
}

func Errorf(c *gin.Context, code ErrCode, format string, values ...interface{}) {
	resp := ErrorResponse{
		Code:    code,
		Message: code.String(),
		Detail:  fmt.Sprintf(format, values...),
	}

	c.JSON(int(code), resp)
}

func ErrorParams(c *gin.Context, err error) {
	Error(c, Err_Invalid_Argument, validator.Translate(err))
}

func OK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"detail":  data,
	})
}

