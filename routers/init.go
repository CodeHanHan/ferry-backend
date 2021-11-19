package routers

import (
	"github.com/CodeHanHan/ferry-backend/middleware"
	"github.com/CodeHanHan/ferry-backend/routers/ping"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	middleware.InitMiddleware(r)

	InitSysRouter(r)

	return r
}

func InitSysRouter(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("")

	ping.PingRouter(g)

	return g
}
