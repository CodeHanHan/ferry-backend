package user

import (
	"github.com/CodeHanHan/ferry-backend/apis/user"
	"github.com/gin-gonic/gin"
)

func RegisterUserRouter(g *gin.RouterGroup, mdw ...gin.HandlerFunc) {
	userGroup := g.Group("/user").Use(mdw...)
	{
		userGroup.GET("/me", user.Profile)
		userGroup.POST("", user.CreateUser)
		userGroup.DELETE("", user.DeleteUser)
		userGroup.PATCH("/updateuser", user.UpdateUser)
		userGroup.PATCH("/adminupdateuser", user.AdminUpdateUser)
		userGroup.PATCH("/changepassword", user.ChangePassword)
		userGroup.POST("/upload", user.UploadAvatar)
	}
}
