package db

import (
	"github.com/CodeHanHan/ferry-backend/pkg/pi"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Store *gorm.DB

func SetUp() {
	Store = pi.Global.Mysql
}
