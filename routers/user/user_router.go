package user

import (
	"github.com/CodeHanHan/ferry-backend/apis/user"
	"github.com/gin-gonic/gin"
)

func RegisterUserRouter(g *gin.RouterGroup, authMiddleware gin.HandlerFunc) {
	pingGroup := g.Group("/user").Use(authMiddleware)
	{
		pingGroup.GET("/me", user.Profile)
	}
}
