package system

import (
	"github.com/CodeHanHan/ferry-backend/apis/system"
	"github.com/gin-gonic/gin"
)

func RegisterPageRouter(g *gin.RouterGroup, mdw ...gin.HandlerFunc) {
	listGroup := g.Use(mdw...)
	{
		listGroup.GET("/menulist", system.GetMenuList)
		listGroup.GET("/rolelist", system.ListRoles)
	}
}

func RegisterRoleMenuRouter(g *gin.RouterGroup, mdw ...gin.HandlerFunc) {
	rolemenuGroup := g.Group("/rolemenu", mdw...)
	{
		rolemenuGroup.GET("", system.GetRoleMenu)
		rolemenuGroup.DELETE("/:id", system.DeleteRoleMenu)
	}
}

func RegisterMenuRouter(g *gin.RouterGroup, mdw ...gin.HandlerFunc) {
	m := g.Group("/menu").Use(mdw...)
	{
		m.POST("", system.CreateMenu)
		m.PUT("/:id", system.UpdateMenu)
		m.DELETE("/:id", system.DeleteMenu)
		m.GET("", system.GetMenu)
	}
}

func RegisterRoleRouter(g *gin.RouterGroup, mdw ...gin.HandlerFunc) {
	roleGroup := g.Group("/role", mdw...)
	{
		roleGroup.POST("", system.CreateRole)
		roleGroup.DELETE("/:role_id", system.DeleteRole)
		roleGroup.GET("", system.ListRoles)
		roleGroup.GET("/:role_id", system.GetRole)
		roleGroup.PUT("", system.UpdateRole)
	}
}
