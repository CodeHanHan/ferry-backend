package ping

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/CodeHanHan/ferry-backend/db/query/ping"
	modelPing "github.com/CodeHanHan/ferry-backend/models/ping"
	"github.com/CodeHanHan/ferry-backend/pkg/app"
	"github.com/CodeHanHan/ferry-backend/pkg/logger"
)

func HandlePing(c *gin.Context) {
	message := c.Query("message")

	reply := fmt.Sprintf("%s, too", message)

	record := modelPing.NewPingRecord(message, reply)

	if err := ping.CreatePingRecord(c, record); err != nil {
		logger.Error(c, err.Error())
		app.Error(c, http.StatusInternalServerError, "创建记录失败: %v", err.Error())
		return
	}

	app.OK(c, reply)
}
