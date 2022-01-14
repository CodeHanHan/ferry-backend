package middleware

import (
	"github.com/gin-gonic/gin"
)

func InitMiddleware(r *gin.Engine) {
	r.Use(Logger())
	r.Use(gin.Recovery())
	r.Use(Options)
	r.Use(NoCache)
	r.Use(Secure)
}
