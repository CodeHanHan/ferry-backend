package middleware

import (
	"github.com/gin-gonic/gin"
)

func InitMiddleware(r *gin.Engine) {
	r.Use(Logger())
	// gin.Logger()
	r.Use(gin.Recovery())
}