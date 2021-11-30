package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	userApis "github.com/CodeHanHan/ferry-backend/apis/user"
	_ "github.com/CodeHanHan/ferry-backend/docs"
	"github.com/CodeHanHan/ferry-backend/middleware"
	"github.com/CodeHanHan/ferry-backend/pkg/pi"
	"github.com/CodeHanHan/ferry-backend/routers/ping"
	"github.com/CodeHanHan/ferry-backend/routers/user"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	middleware.InitMiddleware(r)

	v1 := r.Group("/api/v1")

	InitSwaggerRouter(v1)

	InitNoCheckRouter(v1)

	authMiddleware := middleware.AuthMiddleware(pi.Global.TokenMaker)

	InitAuthSysRouter(v1, authMiddleware)

	return r
}

func InitSwaggerRouter(g *gin.RouterGroup) {
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func InitNoCheckRouter(g *gin.RouterGroup) {
	ping.RegisterPingRouter(g)
}

func InitAuthSysRouter(r *gin.RouterGroup, authMiddleware gin.HandlerFunc) *gin.RouterGroup {
	g := r.Group("")
	g.GET("/login", userApis.Login)

	user.RegisterUserRouter(g, authMiddleware)

	return g
}
