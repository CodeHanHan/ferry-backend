package user

import "time"

// get: /login
type LoginRequest struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	Id       string `json:"uuid" form:"uuid" binding:"required"`
	Code     string `json:"code" form:"code" binding:"required"`
	//LoginType  string `json:"logintype" form:"logintype"`
	//RememberMe bool   `json:"rememberme" form:"rememberme"`
}

type LoginResponse struct {
	Code     int       `json:"code" form:"code"`
	Duration int64     `json:"duration"`
	Expire   time.Time `json:"expire" form:"expire"`
	Token    string    `json:"token" form:"token"`
}

// get: /logintest
type LoginTestRequest struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type LoginTestResponse struct {
	Duration int64  `json:"duration"`
	Token    string `json:"token" form:"token"`
}

// get: /user/me
type ProfileRequest struct {
}

type ProfileResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

// post: /user
type CreateUserRequest struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	Role     string `json:"role" form:"role" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email"`
}

type CreateUserResponse struct {
	ID string `json:"id"`
}

// delete: /user
type DeleteUserRequest struct {
	ID string `json:"id" form:"id" binding:"required"`
}

type DeleteUserResponse struct {
	Result string `json:"result"`
}

type UpdateUserRequest struct {
	Nickname string `json:"nickname" form:"nickname"`
	Password string `json:"password" form:"password"`
	Email    string `json:"email" form:"email" binding:"email"`
}
type AdminUpdateUserRequest struct {
	UpdateUserRequest
	Username string `json:"username" form:"username" binding:"required"`
}

type UpdateUserResponse struct {
	Result string `json:"result"`
}

type ChangePwdRequest struct {
	OldPassword string `json:"oldpassword" form:"oldpassword" binding:"required"`
	NewPassword string `json:"newpassword" form:"newpassword" binding:"required"`
}
