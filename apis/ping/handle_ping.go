package ping

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/CodeHanHan/ferry-backend/db/query/ping"
	modelPing "github.com/CodeHanHan/ferry-backend/models/ping"
	"github.com/CodeHanHan/ferry-backend/pkg/logger"
)

func HandlePing(c *gin.Context) {
	message := c.Query("message")

	reply := fmt.Sprintf("%s, too", message)

	record := modelPing.NewPingRecord(message, reply)

	if err := ping.CreatePingRecord(record); err != nil {
		logger.Error(err.Error())
		c.String(http.StatusInternalServerError, err.Error()) // FIXME
	}

	c.String(http.StatusOK, reply)
}
