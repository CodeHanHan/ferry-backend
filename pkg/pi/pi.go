package pi

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"

	dbLogger "github.com/CodeHanHan/ferry-backend/db/logger"
	"github.com/CodeHanHan/ferry-backend/pkg/config"
	"github.com/CodeHanHan/ferry-backend/pkg/logger"
	"github.com/CodeHanHan/ferry-backend/pkg/token"
)

type Pi struct {
	Cfg        *config.Config
	Mysql      *gorm.DB
	TokenMaker token.Maker
}

var Global *Pi

func SetUp() {
	Global = &Pi{}

	config, err := config.LoadConfig()
	if err != nil {
		logger.Critical(context.Background(), "load config failed: %v", err)
		panic(err)
	}
	Global.Cfg = config
	Global.OpenMysql()
	Global.SetUpTokenMaker()
}

func (p *Pi) OpenMysql() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=%v",
		p.Cfg.Database.DBUser,
		p.Cfg.Database.DBPassword,
		p.Cfg.Database.DBHost,
		p.Cfg.Database.DBPort,
		p.Cfg.Database.DBName,
		p.Cfg.Database.ParseTime,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: dbLogger.
			NewGormLogger(time.Millisecond * 500).
			LogMode(gormLogger.LogLevel(p.Cfg.Database.LoggerLevel)),
	})
	if err != nil {
		logger.Critical(context.Background(), "failed to connect to database: %v", err)
		panic(err)
	}

	p.Mysql = db
}

func (p *Pi) SetUpTokenMaker() {
	tokenMaker, err := token.NewJWTMaker(p.Cfg.Jwt.Secret)
	if err != nil {
		logger.Critical(context.Background(), "failed to create token maker: %v", err)
		panic(err)
	}
	p.TokenMaker = tokenMaker
}
