package ping

import (
	"github.com/CodeHanHan/ferry-backend/apis"
	"github.com/gin-gonic/gin"
)

func PingRouter(r *gin.RouterGroup) {
	r.POST("/ping", apis.HandlePing)
}
