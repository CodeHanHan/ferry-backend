package form

// get: /login
type LoginRequest struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	Id       string `json:"id" form:"id" binding:"required"`
	Code     string `json:"code" form:"code" binding:"required"`
}

type LoginResponse struct {
	Duration int64  `json:"duration"`
	Token    string `json:"token" form:"token"`
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
