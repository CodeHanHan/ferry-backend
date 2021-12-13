package ping

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/CodeHanHan/ferry-backend/db/query/ping"
	modelPing "github.com/CodeHanHan/ferry-backend/models/ping"
	"github.com/CodeHanHan/ferry-backend/pkg/app"
	"github.com/CodeHanHan/ferry-backend/pkg/form"
	"github.com/CodeHanHan/ferry-backend/pkg/logger"
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
// @Router /ping [post]
func Ping(c *gin.Context) {
	// 1. 验证参数
	var req form.PingRequest
	if err := c.ShouldBind(&req); err != nil {
		logger.ErrorParams(c, err)
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
		logger.Error(c, "创建记录失败: %v", err)
		app.InternalServerError(c)
		return
	}

	// 5. 返回信息
	app.OK(c, record)
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
// @Router /ping [get]
func ListPing(c *gin.Context) {
	var req form.ListPingRequest
	if err := c.ShouldBind(&req); err != nil {
		logger.ErrorParams(c, err)
		app.ErrorParams(c, err)
		return
	}

	records, err := ping.ListPingRecords(c, req.Offset, req.Limit)
	if err != nil {
		logger.Error(c, "查询数据库失败: %v", err)
		app.InternalServerError(c)
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
// @Router /ping [delete]
func DeletePing(c *gin.Context) {
	var req form.DeletePingRequest
	if err := c.ShouldBind(&req); err != nil {
		app.ErrorParams(c, err)
		return
	}

	if err := ping.DeletePingRecord(c, req.PingID); err != nil {
		app.InternalServerError(c)
		return
	}

	app.OK(c, nil)
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
// @Router /ping [put]
func UpdatePing(c *gin.Context) {
	// 1. 校验参数
	var req form.UpdatePingRequest
	if err := c.ShouldBind(&req); err != nil {
		app.ErrorParams(c, err)
		return
	}

	// 2. 查询要更改的记录值是否存在
	record, err := ping.GetPingRecordByPingID(c, req.PingID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			app.Errorf(c, app.Err_Not_found, "更新失败，未找到该记录值: %s", req.PingID)
			return
		}
		app.InternalServerError(c)
		return
	}

	// 3. update
	reply := fmt.Sprintf("%s, too", req.UpdateMessage)
	record.Message = req.UpdateMessage
	record.Reply = reply
	if err := ping.UpdatePingRecord(c, record); err != nil {
		app.InternalServerError(c)
		return
	}

	// 4. 返回
	app.OK(c, nil)
}
