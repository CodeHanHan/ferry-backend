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

// Ping godoc
// @Summary 测试服务是否正常启动
// @Description 接收一个字符串，返回这个字符串加上", too"后缀
// @Tags ping
// @ID ping
// @Param message query string true "any string"
// @Success 200 {string} string
// @Accept  json
// @Produce  json
// @Router /ping/create [post]
func Ping(c *gin.Context) {
	// 1. 验证参数
	var req form.PingRequest
	if err := c.ShouldBind(&req); err != nil {
		app.ErrorParams(c, err)
		return
	}

	// 2. 获取参数
	message := req.Message

	// message := c.Query("message")

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

// ListPing godoc
// @Summary 分页获取ping的记录信息
// @Description 接收偏移和限制量，返回对应的ping记录
// @Tags ping
// @ID list-ping
// @Param offset query int true "偏移量"
// @Param limit query int true "每页记录数"
// @Success 200 {object} []modelPing.PingRecord
// @Accept  json
// @Produce  json
// @Router /ping/list [get]
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

// DeletePing godoc
// @Summary 删除记录信息
// @Description 接收主键PingID,根据PingID删除该条记录
// @Tags ping
// @ID delete_ping
// @Param ping_id query string true "any string"
// @Success 200 {string} string
// @Accept  json
// @Produce  json
// @Router /ping/delete [delete]
func DeletePing(c *gin.Context) {
	var req form.DeletePingRequest
	if err := c.ShouldBind(&req); err != nil {
		app.ErrorParams(c, err)
		return
	}

	if err := ping.DeletePingRecord(c, req.PingID); err != nil {
		app.Error(c, err, http.StatusBadRequest, "删除失败")
		return
	}

	app.OK(c, "success")
}

// UpdatePing godoc
// @Summary 更新记录信息
// @Description 接收主键PingID,根据PingID更新该条记录的message
// @Tags ping
// @ID update_ping
// @Param ping_id query string true "any string"
// @Param updatemessage query string true "any string"
// @Success 200 {string} string
// @Accept  json
// @Produce  json
// @Router /ping/update [put]
func UpdatePing(c *gin.Context) {

	var req form.UpdatePingRequest
	if err := c.ShouldBind(&req); err != nil {
		app.ErrorParams(c, err)
		return
	}
	reply := fmt.Sprintf("%s, too", req.UpdateMessage)
	if err := ping.UpdatePingRecord(c, req.PingID, req.UpdateMessage, reply); err != nil {
		app.Error(c, err, http.StatusBadRequest, "更新失败")
		return
	}

	app.OK(c, "更新成功")
}
