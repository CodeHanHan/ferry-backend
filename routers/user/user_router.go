package user

import (
	"github.com/CodeHanHan/ferry-backend/apis/user"
	"github.com/gin-gonic/gin"
)

func RegisterUserRouter(g *gin.RouterGroup, mdw ...gin.HandlerFunc) {
	pingGroup := g.Group("/user").Use(mdw...)
	{
		pingGroup.GET("/me", user.Profile)
		pingGroup.POST("", user.CreateSysUser)
		pingGroup.DELETE("", user.DeleteSysUser)
	}
}
