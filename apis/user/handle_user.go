package user

import (
	"fmt"
	"net/http"
	"time"

	"github.com/CodeHanHan/ferry-backend/db/query/user"
	modelUsers "github.com/CodeHanHan/ferry-backend/models/users"
	"github.com/CodeHanHan/ferry-backend/pkg/app"
	"github.com/CodeHanHan/ferry-backend/pkg/form"
	"github.com/CodeHanHan/ferry-backend/pkg/logger"
	"github.com/CodeHanHan/ferry-backend/pkg/pi"
	"github.com/CodeHanHan/ferry-backend/pkg/token"
	"github.com/CodeHanHan/ferry-backend/utils/password"
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
	var req form.LoginRequest
	if err := c.ShouldBind(&req); err != nil {
		logger.Error(c, "参数验证失败: %v", err)
		app.ErrorParams(c, err)
	}

	username := req.Username
	pwd := req.Password

	query, err := user.GetByUserName(c, username)
	if err != nil {
		app.Error(c, http.StatusBadRequest, "该用户不存在")
		return
	}

	if err := password.CheckPassword(pwd, query.Password); err != nil {
		app.Error(c, http.StatusBadRequest, "密码输入错误")
		return
	}

	jwtToken, err := pi.Global.TokenMaker.CreateToken(username, query.Role, time.Hour)
	if err != nil {
		logger.Error(c, "生成token失败: %v", err)
		app.InternalServerError(c)
		return
	}

	app.OK(c, form.LoginResponse{
		Duration: time.Hour.Microseconds(),
		Token:    jwtToken,
	})
}

// Profile godoc
// @Summary 查看个人信息
// @Description 用户查看个人信息
// @Tags user
// @ID user-me
// @Success 200 {object} form.ProfileResponse
// @Accept  json
// @Produce  json
// @Router /user/me [get]
// @Security BearerAuth
func Profile(c *gin.Context) {
	payload, err := token.GetPayload(c)
	if err != nil {
		logger.Error(c, "获取payload失败")
		app.InternalServerError(c)
	}

	username := payload.Username
	email := fmt.Sprintf("%s@ferry.com", username)

	if username == "admin" {
		app.OK(c, &form.ProfileResponse{Username: username, Email: email})
		return
	}

	app.OK(c, "非admin用户")
}

// Register godoc
// @Summary 创建用户信息
// @Description 管理员创建用户个人信息
// @Tags user
// @ID user-insertsysuser
// @Param username query string true "用户名"
// @Param password query string true "密码"
// @Param role query string true "角色"
// @Param email query string true "邮箱"
// @Success 200 {object} form.InsertSysUserRequest
// @Accept  json
// @Produce  json
// @Router /user/insertsysuser [post]
// @Security BearerAuth
func InsertSysUser(c *gin.Context) {
	var req form.InsertSysUserRequest
	if err := c.ShouldBind(&req); err != nil {
		logger.Error(c, "参数验证失败: %v", err)
		app.ErrorParams(c, err)
		return
	}

	username := req.Username
	pwd := req.Password
	role := req.Role
	email := req.Email

	crypto_pwd, err := password.HashPassword(pwd)
	if err != nil {
		logger.Error(c, "密码加密失败")
		app.InternalServerError(c)
		return
	}
	record := modelUsers.NewUsersTable(username, crypto_pwd, role, email)
	if err := user.CreateUserRecord(c, record); err != nil {
		logger.Error(c, "创建记录失败: %v", err)
		app.InternalServerError(c)
		return
	}
	app.OK(c, "创建成功")
}
