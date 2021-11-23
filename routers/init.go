package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	_ "github.com/CodeHanHan/ferry-backend/docs"
	"github.com/CodeHanHan/ferry-backend/middleware"
	"github.com/CodeHanHan/ferry-backend/routers/ping"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	middleware.InitMiddleware(r)

	v1 := r.Group("/api/v1")

	InitSysRouter(v1)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}

func InitSysRouter(r *gin.RouterGroup) *gin.RouterGroup {
	g := r.Group("")

	ping.RegisterPingRouter(g)

	return g
}
