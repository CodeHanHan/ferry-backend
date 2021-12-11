package ping

import (
	"github.com/CodeHanHan/ferry-backend/apis/ping"
	"github.com/gin-gonic/gin"
)

func RegisterPingRouter(g *gin.RouterGroup) {
	pingGroup := g.Group("/ping")
	{
		pingGroup.POST("", ping.Ping)
		pingGroup.GET("", ping.ListPing)
		pingGroup.DELETE("", ping.DeletePing)
		pingGroup.PUT("", ping.UpdatePing)
	}
}
