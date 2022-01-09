package dept

import (
	"github.com/CodeHanHan/ferry-backend/apis/dept"
	"github.com/gin-gonic/gin"
)

func RegisterDeptRouter(g *gin.RouterGroup, mdw ...gin.HandlerFunc) {
	deptGroup := g.Group("/dept", mdw...)
	{
		deptGroup.POST("", dept.CreateDept)
		deptGroup.DELETE("/:dept_id", dept.DeleteDept)
		deptGroup.GET("", dept.ListDpet)
		deptGroup.GET("/:dept_id", dept.GetDept)
		deptGroup.PUT("", dept.UpdateDept)
	}
}
