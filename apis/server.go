package apis

import (
	"github.com/CodeHanHan/ferry-backend/db"
	"github.com/gin-gonic/gin"
)

type server struct {
	db     db.Database
	router *gin.Engine
}

func NewServer(db db.Database) *server {
	return &server{
		db:     db,
		router: &gin.Engine{}, // FIXME
	}
}
