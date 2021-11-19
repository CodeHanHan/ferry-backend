package main

import (
	"log"
	"net/http"

	"github.com/CodeHanHan/ferry-backend/db"
	"github.com/CodeHanHan/ferry-backend/pkg/pi"
	"github.com/CodeHanHan/ferry-backend/routers"
)

func main() {
	// time.Sleep(time.Second * 10)

	pi.SetUp()

	db.SetUp()

	r := routers.InitRouter()

	// gin.SetMode(gin.DebugMode)

	server := &http.Server{
		Addr:    pi.Global.Cfg.Application.ServerAddress,
		Handler: r,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Printf("Server error: %v", err)
	}
}
