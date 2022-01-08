package role

import (
	"github.com/CodeHanHan/ferry-backend/apis/role"
	"github.com/gin-gonic/gin"
)

func RegisterRoleRouter(g *gin.RouterGroup, mdw ...gin.HandlerFunc) {
	roleGroup := g.Group("/role", mdw...)
	{
		roleGroup.POST("", role.CreateRole)
		roleGroup.DELETE("/:role_id", role.DeleteRole)
		roleGroup.GET("", role.ListRoles)
	}
}
