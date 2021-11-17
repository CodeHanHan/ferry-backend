package main

import (
	"log"
	"net/http"

	"github.com/CodeHanHan/ferry-backend/pkg/config"
	"github.com/CodeHanHan/ferry-backend/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("load config failed: %v", err)
	}
	r := routers.InitRouter()

	gin.SetMode(gin.DebugMode)

	server := &http.Server{
		Addr:    config.Application.ServerAddress,
		Handler: r,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Printf("Server error: %v", err)
	}
}
