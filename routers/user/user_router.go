package user

import (
	"github.com/CodeHanHan/ferry-backend/pkg/jwtauth"
	"github.com/gin-gonic/gin"
)

func RegisterUserRouter(g *gin.RouterGroup, authMiddleware *jwtauth.GinJWTMiddleware) {
	pingGroup := g.Group("/user")
	{
		pingGroup.POST("/login", authMiddleware.LoginHandler)
	}
}
