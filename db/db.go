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

type Filter struct {
	Params map[string]interface{}
}

func NewFilter() *Filter {
	return &Filter{
		Params: make(map[string]interface{}),
	}
}

func (f *Filter) Set(k string, v interface{}) *Filter {
	f.Params[k] = v
	return f
}
