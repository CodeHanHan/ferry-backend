package ping

import (
	"github.com/CodeHanHan/ferry-backend/apis/ping"
	"github.com/gin-gonic/gin"
)

func RegisterPingRouter(g *gin.RouterGroup) {
	pingGroup := g.Group("/ping")
	{
		pingGroup.POST("/create", ping.Ping)
		pingGroup.GET("/list", ping.ListPing)
		// delete
		pingGroup.DELETE("/delete", ping.DeletePing)
		// update

	}
}
