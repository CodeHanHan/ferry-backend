package menu

import (
	"github.com/CodeHanHan/ferry-backend/apis/menu"
	"github.com/gin-gonic/gin"
)

func RegisterMenuRouter(g *gin.RouterGroup, mdw ...gin.HandlerFunc) {
	m := g.Group("/menu").Use(mdw...)
	{
		m.POST("", menu.CreateMenu)
	}
}
