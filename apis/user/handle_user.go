package user

import (
	"time"

	"github.com/CodeHanHan/ferry-backend/db"
	"github.com/CodeHanHan/ferry-backend/db/query/user"
	modelUser "github.com/CodeHanHan/ferry-backend/models/user"
	"github.com/CodeHanHan/ferry-backend/pkg/app"
	formUser "github.com/CodeHanHan/ferry-backend/pkg/form/user"
	"github.com/CodeHanHan/ferry-backend/pkg/logger"
	"github.com/CodeHanHan/ferry-backend/pkg/pi"
	"github.com/CodeHanHan/ferry-backend/pkg/sender"
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
// @Param id query string true "验证码id"
// @Param code query string true "验证码内容"
// @Success 200 {object} formUser.LoginResponse
// @Failure 500 {object} app.ErrResponse
// @Failure 400 {object} app.ErrResponse
// @Produce  json
// @Router /login [get]
func Login(c *gin.Context) {
	var req formUser.LoginRequest
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

	app.OK(c, formUser.LoginResponse{
		Duration: time.Hour.Microseconds(),
		Token:    jwtToken,
	})
}

// LoginTest godoc
// @Summary 用户名密码登录(测试用)
// @Description 获取token
// @Tags user
// @ID user-logintest
// @Param username query string true "用户名"
// @Param password query string true "密码"
// @Success 200 {object} formUser.LoginResponse
// @Failure 500 {object} app.ErrResponse
// @Failure 400 {object} app.ErrResponse
// @Produce  json
// @Router /logintest [get]
func LoginTest(c *gin.Context) {
	var req formUser.LoginTestRequest
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

	app.OK(c, formUser.LoginResponse{
		Duration: time.Hour.Microseconds(),
		Token:    jwtToken,
	})
}

// Profile godoc
// @Summary 查看个人信息
// @Description 用户查看个人信息
// @Tags user
// @ID user-me
// @Success 200 {object} formUser.ProfileResponse
// @Failure 500 {object} app.ErrResponse
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

	app.OK(c, formUser.ProfileResponse{
		Username: userRecord.Username,
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
// @Success 200 {object} form.CreateUserResponse
// @Failure 500 {object} app.ErrResponse
// @Failure 400 {object} app.ErrResponse
// @Produce  json
// @Router /user [post]
// @Security BearerAuth
func CreateUser(c *gin.Context) {
	var req formUser.CreateUserRequest
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
	record := modelUser.NewUserTable(username, crypto_pwd, role, email)
	if err := user.CreateUserRecord(c, record); err != nil {
		logger.Error(c, "创建记录失败: %v", err)
		app.InternalServerError(c)
		return
	}
	app.OK(c, formUser.CreateUserResponse{
		ID: record.ID,
	})
}

// Delete godoc
// @Summary 删除用户信息
// @Description 管理员删除用户个人信息
// @Tags user
// @ID user-deletesysuser
// @Param id query string true "用户ID"
// @Success 200 {object} formUser.DeleteUserResponse
// @Failure 500 {object} app.ErrResponse
// @Failure 400 {object} app.ErrResponse
// @Produce  json
// @Router /user [delete]
// @Security BearerAuth
func DeleteUser(c *gin.Context) {
	var req formUser.DeleteUserRequest
	if err := c.ShouldBind(&req); err != nil {
		logger.Error(c, "获取信息失败")
		app.ErrorParams(c, err)
		return
	}

	uid := req.ID
	if uid == "0" {
		logger.Error(c, "非法删除")
		app.Error(c, app.Err_Permission_Denied, "删除权限不够")
		return
	}

	if err := user.DeleteSysUser(c, uid); err != nil {
		logger.Error(c, "删除失败: %v", err)
		app.InternalServerError(c)
		return
	}

	app.OK(c, formUser.DeleteUserResponse{
		Result: "success",
	})
}

// Upadate godoc
// @Summary 更新用户信息
// @Description 用户更新个人信息
// @Tags user
// @ID user-updateuser
// @Param nickname query string false "昵称"
// @Param email query string false "邮箱"
// @Success 200 {object} formUser.UpdateUserResponse
// @Failure 500 {object} app.ErrResponse
// @Failure 400 {object} app.ErrResponse
// @Produce  json
// @Router /user/updateuser [patch]
// @Security BearerAuth
func UpdateUser(c *gin.Context) {
	sendername, _, err := sender.GetSender(c)
	if err != nil {
		app.InternalServerError(c)
		return
	}

	var req formUser.UpdateUserRequest
	if err := c.ShouldBind(&req); err != nil {
		logger.Error(c, "获取信息失败: %v", err)
		app.ErrorParams(c, err)
		return
	}

	nickname := req.Nickname
	email := req.Email
	filter1 := db.NewFilter().Set("username", sendername)
	filter2 := db.NewFilter().Set("nickname", nickname).Set("email", email)

	if err := user.UpdateUserRecord(c, filter1, filter2); err != nil {
		logger.Error(c, "更新信息失败: %v", err)
		app.InternalServerError(c)
		return
	}

	app.OK(c, formUser.UpdateUserResponse{
		Result: "success",
	})
}

// AdminUpadate godoc
// @Summary 管理员更新用户信息
// @Description 管理员更新个人信息
// @Tags user
// @ID user-adminupdateuser
// @Param username query string true "Username"
// @Param nickname query string false "昵称"
// @Param password query string false "密码"
// @Param email query string false "邮箱"
// @Success 200 {object} formUser.UpdateUserResponse
// @Failure 500 {object} app.ErrResponse
// @Failure 400 {object} app.ErrResponse
// @Produce  json
// @Router /user/adminupdateuser [patch]
// @Security BearerAuth
func AdminUpdateUser(c *gin.Context) {
	var req formUser.AdminUpdateUserRequest
	if err := c.ShouldBind(&req); err != nil {
		logger.Error(c, "获取信息失败: %v", err)
		app.ErrorParams(c, err)
		return
	}
	username := req.Username
	nickname := req.Nickname
	pwd := req.Password
	email := req.Email

	crypto_pwd, err := password.HashPassword(pwd)
	if err != nil {
		logger.Error(c, "密码加密失败: %v", err)
		app.InternalServerError(c)
		return
	}

	filter1 := db.NewFilter().Set("username", username)
	filter2 := db.NewFilter().Set("nickname", nickname).Set("password", crypto_pwd).Set("email", email)

	if err := user.UpdateUserRecord(c, filter1, filter2); err != nil {
		logger.Error(c, "更新信息失败: %v", err)
		app.InternalServerError(c)
		return

	}

	app.OK(c, formUser.UpdateUserResponse{
		Result: "success",
	})
}

// Changepassword godoc
// @Summary 用户更新密码
// @Description 用户更新密码
// @Tags user
// @ID user-changepassword
// @Param oldpassword query string true "旧密码"
// @Param newpassword query string true "新密码"
// @Success 200 {object} formUser.UpdateUserResponse
// @Failure 500 {object} app.ErrResponse
// @Failure 400 {object} app.ErrResponse
// @Produce  json
// @Router /user/changepassword [patch]
// @Security BearerAuth
func ChangePassword(c *gin.Context) {
	var req formUser.ChangePwdRequest
	if err := c.ShouldBind(&req); err != nil {
		logger.Error(c, "获取信息失败: %v", err)
		app.ErrorParams(c, err)
		return
	}
	newPwd := req.NewPassword
	oldPwd := req.OldPassword

	sender, _, err := sender.GetSender(c)
	if err != nil {
		app.InternalServerError(c)
		return
	}

	filter1 := db.NewFilter().Set("username", sender)
	userRecord, err := user.GetByUserName(c, filter1)
	if err != nil {
		app.InternalServerError(c)
		return
	}

	if err := password.CheckPassword(oldPwd, userRecord.Password); err != nil { // check用户输入的原密码和数据库中的密码
		app.Error(c, app.Err_Permission_Denied, "原始密码输入错误")
		return
	}

	crypto_pwd, err := password.HashPassword(newPwd) //新密码进行密码加密
	if err != nil {
		logger.Error(c, "密码加密失败: %v", err)
		app.InternalServerError(c)
		return
	}

	filter2 := db.NewFilter().Set("password", crypto_pwd)
	if err := user.UpdateUserRecord(c, filter1, filter2); err != nil { //将新密码插入到数据库中
		logger.Error(c, "更新密码失败: %v", err)
		app.InternalServerError(c)
	}

	app.OK(c, formUser.UpdateUserResponse{
		Result: "success",
	})
}
