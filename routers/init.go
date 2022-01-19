package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	userApis "github.com/CodeHanHan/ferry-backend/apis/user"
	_ "github.com/CodeHanHan/ferry-backend/docs"
	"github.com/CodeHanHan/ferry-backend/middleware"
	"github.com/CodeHanHan/ferry-backend/pkg/pi"
	"github.com/CodeHanHan/ferry-backend/routers/dept"
	"github.com/CodeHanHan/ferry-backend/routers/ping"
	"github.com/CodeHanHan/ferry-backend/routers/post"
	"github.com/CodeHanHan/ferry-backend/routers/role"
	"github.com/CodeHanHan/ferry-backend/routers/user"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	middleware.InitMiddleware(r)

	v1 := r.Group("/api/v1")

	InitSwaggerRouter(v1)

	InitNoCheckRouter(v1)

	authMiddleware := middleware.AuthMiddleware(pi.Global.TokenMaker)
	roleMiddleware := middleware.CheckRole()
	senderMiddleware := middleware.Sender()

	InitAuthSysRouter(v1, authMiddleware, roleMiddleware, senderMiddleware)

	return r
}

func InitSwaggerRouter(g *gin.RouterGroup) {
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func InitNoCheckRouter(g *gin.RouterGroup) {
	ping.RegisterPingRouter(g)

	g.GET("/captcha", userApis.Captcha)
	g.POST("/captcha", userApis.VerifyCaptcha)
}

func InitAuthSysRouter(r *gin.RouterGroup, mdw ...gin.HandlerFunc) *gin.RouterGroup {
	g := r.Group("")
	g.POST("/login", userApis.Login)
	g.GET("/logintest", userApis.LoginTest)

	user.RegisterUserRouter(g, mdw...)
	role.RegisterRoleRouter(g, mdw...)
	dept.RegisterDeptRouter(g, mdw...)
	post.RegisterPostRouter(g, mdw...)
	return g
}
