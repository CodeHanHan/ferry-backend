package user

import (
	"github.com/CodeHanHan/ferry-backend/apis/user"
	"github.com/gin-gonic/gin"
)

func RegisterUserRouter(g *gin.RouterGroup, authMiddleware gin.HandlerFunc, roleMiddleware gin.HandlerFunc) {
	pingGroup := g.Group("/user").Use(authMiddleware).Use(roleMiddleware)
	{
		pingGroup.GET("/me", user.Profile)
		pingGroup.POST("/createuser", user.InsertSysUser)
	}
}
