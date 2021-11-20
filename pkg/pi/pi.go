package pi

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	dbLogger "github.com/CodeHanHan/ferry-backend/db/logger"
	"github.com/CodeHanHan/ferry-backend/pkg/config"
)

type Pi struct {
	Cfg   *config.Config
	Mysql *gorm.DB
}

var Global *Pi

func SetUp() {
	Global = &Pi{}

	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("load config failed: %v", err)
	}
	Global.Cfg = config
	Global.OpenMysql()
}

func (p *Pi) OpenMysql() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		p.Cfg.Database.DBUser,
		p.Cfg.Database.DBPassword,
		p.Cfg.Database.DBHost,
		p.Cfg.Database.DBPort,
		p.Cfg.Database.DBName,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: dbLogger.NewGormLogger(time.Millisecond * 500),
	})
	if err != nil {
		log.Fatal(err)
	}

	p.Mysql = db
}
