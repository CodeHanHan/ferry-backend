package user

import (
	"time"

	"github.com/gin-gonic/gin"

	"github.com/CodeHanHan/ferry-backend/db"
	"github.com/CodeHanHan/ferry-backend/db/query/user"
	modelUsers "github.com/CodeHanHan/ferry-backend/models/users"
	"github.com/CodeHanHan/ferry-backend/pkg/app"
	"github.com/CodeHanHan/ferry-backend/pkg/form"
	"github.com/CodeHanHan/ferry-backend/pkg/logger"
	"github.com/CodeHanHan/ferry-backend/pkg/pi"
	"github.com/CodeHanHan/ferry-backend/pkg/sender"
	"github.com/CodeHanHan/ferry-backend/utils/password"
)

// Login godoc
// @Summary 用户名密码登录
// @Description 获取token
// @Tags user
// @ID user-login
// @Param username query string true "用户名"
// @Param password query string true "密码"
// @Param id query string true "验证码id"
// @Param code query string true "验证码内容"
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
	id := req.Id
	code := req.Code

	ok := store.Verify(id, code, true)
	if !ok {
		app.Error(c, app.Err_Permission_Denied, "验证码错误")
		return
	}

	filter := db.NewFilter().Set("username", username)
	query, err := user.GetByUserName(c, filter)
	if err != nil {
		app.Error(c, app.Err_Permission_Denied, "该用户不存在")
		return
	}

	if err := password.CheckPassword(pwd, query.Password); err != nil {
		app.Error(c, app.Err_Permission_Denied, "密码输入错误")
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

// LoginTest godoc
// @Summary 用户名密码登录
// @Description 获取token
// @Tags user
// @ID user-logintest
// @Param username query string true "用户名"
// @Param password query string true "密码"
// @Success 200 {object} form.LoginResponse
// @Accept  json
// @Produce  json
// @Router /logintest [get]
func LoginTest(c *gin.Context) {
	var req form.LoginTestRequest
	if err := c.ShouldBind(&req); err != nil {
		logger.Error(c, "参数验证失败: %v", err)
		app.ErrorParams(c, err)
	}

	username := req.Username
	pwd := req.Password

	filter := db.NewFilter().Set("username", username)
	query, err := user.GetByUserName(c, filter)
	if err != nil {
		app.Error(c, app.Err_Permission_Denied, "该用户不存在")
		return
	}

	if err := password.CheckPassword(pwd, query.Password); err != nil {
		app.Error(c, app.Err_Permission_Denied, "密码输入错误")
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
	sender, _, err := sender.GetSender(c)
	if err != nil {
		app.InternalServerError(c)
		return
	}

	filter := db.NewFilter().Set("username", sender)
	userRecord, err := user.GetByUserName(c, filter)
	if err != nil {
		app.InternalServerError(c)
		return
	}

	app.OK(c, form.ProfileResponse{
		Username: userRecord.UserName,
		Email:    userRecord.Email,
	})
}

// Register godoc
// @Summary 创建用户信息
// @Description 管理员创建用户个人信息
// @Tags user
// @ID user-createsysuser
// @Param username query string true "用户名"
// @Param password query string true "密码"
// @Param role query string true "角色"
// @Param email query string true "邮箱"
// @Success 200 {object} form.CreateSysUserRequest
// @Accept  json
// @Produce  json
// @Router /user [post]
// @Security BearerAuth
func CreateSysUser(c *gin.Context) {
	var req form.CreateSysUserRequest
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

// Delete godoc
// @Summary 删除用户信息
// @Description 管理员删除用户个人信息
// @Tags user
// @ID user-deletesysuser
// @Param id query string true "用户ID"
// @Success 200 {object} form.DeleteSysUserRequest
// @Accept  json
// @Produce  json
// @Router /user [delete]
// @Security BearerAuth
func DeleteSysUser(c *gin.Context) {
	var req form.DeleteSysUserRequest
	if err := c.ShouldBind(&req); err != nil {
		logger.Error(c, "获取信息失败")
		app.ErrorParams(c, err)
		return
	}

	/* senderUsername, _, err := sender.GetSender(c)
	if err != nil {
		app.InternalServerError(c)
		return
	} */

	uid := req.ID
	if uid == "0" {
		logger.Error(c, "非法删除")
		app.Error(c, app.Err_Permission_Denied, "删除权限不够")
		return
	}
	if err := user.DeleteSysUser(c, uid); err != nil {
		logger.Error(c, "删除失败:%v", err)
		app.InternalServerError(c)
		return
	}
	app.OK(c, "删除成功")
}
