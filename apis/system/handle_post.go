package system

import (
	"github.com/CodeHanHan/ferry-backend/models/system"
	"github.com/CodeHanHan/ferry-backend/pkg/app"
	"github.com/CodeHanHan/ferry-backend/pkg/logger"
	"github.com/CodeHanHan/ferry-backend/pkg/sender"
	"github.com/gin-gonic/gin"
)

// @Summary 添加职位
// @Description 获取JSON
// @Tags 职位
// @Accept  application/json
// @Product application/json
// @Param data body system.Post true "data"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /post [post]
// @Security Bearer
func InsertPost(c *gin.Context) {
	var post system.Post
	if err := c.Bind(&post); err != nil {
		logger.Error(c, err.Error())
		app.ErrorParams(c, err)
		return
	}

	sender, _, err := sender.GetSender(c)
	if err != nil {
		app.InternalServerError(c)
		return
	}

	post.CreateBy = sender

	post_, err := post.Create(c)
	if err != nil {
		app.InternalServerError(c)
		return
	}

	app.AdaptOK(c, post_, "")
}
