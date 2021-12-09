package main

import (
	"context"
	"log"
	"net/http"

	"github.com/CodeHanHan/ferry-backend/db"
	"github.com/CodeHanHan/ferry-backend/pkg/logger"
	"github.com/CodeHanHan/ferry-backend/pkg/pi"
	"github.com/CodeHanHan/ferry-backend/routers"
	"github.com/golang-migrate/migrate/v4"
)

// @title Ferry API
// @version v0.0.1
// @description 工单系统

// @host localhost:10000
// @BasePath /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	// time.Sleep(time.Second * 10)

	if err := pi.SetUp(); err != nil {
		panic(err)
	}

	db.SetUp()

	r := routers.InitRouter()

	server := &http.Server{
		Addr:    pi.Global.Cfg.Application.ServerAddress,
		Handler: r,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Printf("Server error: %v", err)
	}
}

func Migrate(source string, dsn string) error {
	m, err := migrate.New(source, dsn)
	if err != nil {
		logger.Error(context.Background(), "Failed to load db source or connection: %v", err)
		return err
	}

	// Migrate all the way up ...
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		logger.Error(context.Background(), "Failed to migrate: %v", err)
		return err
	}

	return nil
}
