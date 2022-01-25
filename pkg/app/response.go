package app

import (
	"fmt"
	"net/http"

	"github.com/CodeHanHan/ferry-backend/pkg/validator"
	"github.com/gin-gonic/gin"
)

type ErrResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Detail  interface{} `json:"details"`
}

func Error(c *gin.Context, code ErrCode, data interface{}) {
	resp := ErrResponse{
		Code:    int(code),
		Message: code.String(),
		Detail:  data,
	}

	c.JSON(int(code), resp)
}

func InternalServerError(c *gin.Context) {
	Error(c, Err_Internal, nil)
}

func Errorf(c *gin.Context, code ErrCode, format string, values ...interface{}) {
	Error(c, code, fmt.Sprintf(format, values...))
}

func ErrorParams(c *gin.Context, err error) {
	Error(c, Err_Invalid_Argument, validator.Translate(err))
}

func OK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

// AdaptOK 用于对接前端(暂时替代OK)
func AdaptOK(c *gin.Context, data interface{}, msg string) {
	var res Response
	res.Data = data
	if msg != "" {
		res.Msg = msg
	}
	c.JSON(http.StatusOK, res.ReturnOK())
}

func Custom(c *gin.Context, h gin.H) {
	c.JSON(http.StatusOK, h)
}

// 分页数据处理
func PageOK(c *gin.Context, result interface{}, count int, pageIndex int, pageSize int, msg string) {
	var res PageResponse
	res.Data.List = result
	res.Data.Count = count
	res.Data.PageIndex = pageIndex
	res.Data.PageSize = pageSize
	if msg != "" {
		res.Msg = msg
	}
	c.JSON(http.StatusOK, res.ReturnOK())
}
