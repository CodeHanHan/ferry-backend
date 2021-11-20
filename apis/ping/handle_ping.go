package ping

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/CodeHanHan/ferry-backend/db/query/ping"
	modelPing "github.com/CodeHanHan/ferry-backend/models/ping"
	"github.com/CodeHanHan/ferry-backend/pkg/app"
	"github.com/CodeHanHan/ferry-backend/pkg/form"
)

func Ping(c *gin.Context) {
	// 1. 验证参数
	var req form.PingRequest
	if err := c.ShouldBind(&req); err != nil {
		app.ErrorParams(c, err)
		return
	}

	// 2. 获取参数
	message := req.Message

	// 3. 逻辑处理，生成回复信息
	reply := fmt.Sprintf("%s, too", message)

	// 4. CRUD
	record := modelPing.NewPingRecord(message, reply)
	if err := ping.CreatePingRecord(c, record); err != nil {
		// 4. 返回前端信息
		app.Error(c, err, http.StatusInternalServerError, "创建记录失败: %v", err.Error())
		return
	}

	// 5. 返回信息
	app.OK(c, reply)
}

func ListPing(c *gin.Context) {
	var req form.ListPingRequest
	if err := c.ShouldBind(&req); err != nil {
		app.ErrorParams(c, err)
		return
	}

	records, err := ping.PagePingRecords(c, req.Offset, req.Limit)
	if err != nil {
		app.Error(c, err, http.StatusBadRequest, "查询失败")
		return
	}

	app.OK(c, records)
}
