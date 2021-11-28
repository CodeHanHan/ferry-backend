package routers

import (
	"context"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	_ "github.com/CodeHanHan/ferry-backend/docs"
	"github.com/CodeHanHan/ferry-backend/middleware"
	"github.com/CodeHanHan/ferry-backend/pkg/jwtauth"
	"github.com/CodeHanHan/ferry-backend/pkg/logger"
	"github.com/CodeHanHan/ferry-backend/routers/ping"
	"github.com/CodeHanHan/ferry-backend/routers/user"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	middleware.InitMiddleware(r)

	authMiddleware, err := middleware.AuthInit()
	if err != nil {
		logger.Error(context.Background(), "init auth middleware failed: %v", err)
		panic(err)
	}

	v1 := r.Group("/api/v1")

	InitSysRouter(v1, authMiddleware)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}

func InitSysRouter(r *gin.RouterGroup, authMiddleware *jwtauth.GinJWTMiddleware) *gin.RouterGroup {
	g := r.Group("")

	ping.RegisterPingRouter(g)
	user.RegisterUserRouter(g, authMiddleware)

	return g
}
