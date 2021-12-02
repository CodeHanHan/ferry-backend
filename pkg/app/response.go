package app

import (
	"net/http"

	"github.com/CodeHanHan/ferry-backend/pkg/validator"
	"github.com/gin-gonic/gin"
)

func InternalServerError(c *gin.Context) {
	c.String(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
}

func ErrorParams(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, validator.Translate(err))
}

func OK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}
