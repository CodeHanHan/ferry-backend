package pi

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/CodeHanHan/ferry-backend/pkg/config"
	"github.com/CodeHanHan/ferry-backend/pkg/mycasbin"
	"github.com/CodeHanHan/ferry-backend/pkg/token"
	"github.com/CodeHanHan/ferry-backend/pkg/validator"
	"github.com/CodeHanHan/ferry-backend/pkg/xmysql"
	"github.com/casbin/casbin/v2"
)

type Pi struct {
	Cfg        *config.Config
	Mysql      *gorm.DB
	TokenMaker token.Maker
	Casbin     *casbin.Enforcer
}

var Global *Pi

func SetUp() error {
	Global = &Pi{}

	// load config
	if err := config.SetUp(func(cfg *config.Config) {
		Global.Cfg = cfg
	}); err != nil {
		return err
	}

	// setup mysql
	if err := xmysql.SetUp(func(db *gorm.DB) {
		Global.Mysql = db
	}, fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=%v",
		Global.Cfg.Database.DBUser,
		Global.Cfg.Database.DBPassword,
		Global.Cfg.Database.DBHost,
		Global.Cfg.Database.DBPort,
		Global.Cfg.Database.DBName,
		Global.Cfg.Database.ParseTime,
	), Global.Cfg.Database.LoggerLevel); err != nil {
		return err
	}

	// setup token maker
	if err := token.SetUp(func(maker token.Maker) {
		Global.TokenMaker = maker
	}, Global.Cfg.Jwt.Secret); err != nil {
		return nil
	}

	// setup permissions check tools
	if err := mycasbin.SetUp(func(e *casbin.Enforcer) {
		Global.Casbin = e
	}, Global.Mysql); err != nil {
		return err
	}

	if err := validator.SetUp(); err != nil {
		return err
	}

	return nil
}
