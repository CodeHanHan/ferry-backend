package user

import (
	"fmt"
	"net/http"
	"time"

	"github.com/CodeHanHan/ferry-backend/pkg/app"
	"github.com/CodeHanHan/ferry-backend/pkg/form"
	"github.com/CodeHanHan/ferry-backend/pkg/logger"
	"github.com/CodeHanHan/ferry-backend/pkg/pi"
	"github.com/CodeHanHan/ferry-backend/pkg/token"
	"github.com/gin-gonic/gin"
)

// Login godoc
// @Summary 用户名密码登录
// @Description 获取token
// @Tags user
// @ID user-login
// @Param username query string true "用户名"
// @Param password query string true "密码"
// @Success 200 {object} form.LoginResponse
// @Accept  json
// @Produce  json
// @Router /login [get]
func Login(c *gin.Context) {
	var loginReq form.LoginRequest
	if err := c.ShouldBind(&loginReq); err != nil {
		logger.Error(c, "参数验证失败: %v", err)
		app.ErrorParams(c, err)
	}

	username := loginReq.Username
	password := loginReq.Password
	if username == "admin" && password == "admin" { // FIXME 硬编码， 改成从数据库查询，验证密码
		jwtToken, err := pi.Global.TokenMaker.CreateToken(username, time.Hour)
		if err != nil {
			app.Error(c, err, http.StatusBadRequest, "生成token失败")
			return
		}

		app.OK(c, form.LoginResponse{
			Duration: time.Hour.Microseconds(),
			Token:    jwtToken,
		})
		return
	}

	app.OK(c, "用户名密码错误")
}

// Profile godoc
// @Summary 查看个人信息
// @Description 用户查看个人信息
// @Tags user
// @ID user-me
// @Success 200 {object} string
// @Accept  json
// @Produce  json
// @Router /user/me [get]
// @Security BearerAuth
func Profile(c *gin.Context) {
	payload, err := token.GetPayload(c)
	if err != nil {
		logger.Error(c, "获取payload失败")
		app.Error(c, err, http.StatusBadRequest, "")
	}

	username := payload.Username
	email := fmt.Sprintf("%s@ferry.com", username)

	if username == "admin" {
		app.OK(c, &form.ProfileResponse{Username: username, Email: email})
		return
	}

	app.OK(c, "非admin用户")
}
