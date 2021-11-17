package middleware

import "github.com/gin-gonic/gin"

func InitMiddleware(r *gin.Engine) {
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
}
