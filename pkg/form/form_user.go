package form

type LoginRequest struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type LoginResponse struct {
	Duration int64  `json:"duration"`
	Token    string `json:"token" form:"token"`
}

type ProfileRequest struct {
}

type ProfileResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type RegisterRequest struct {
	ID string `json:"id" form:"id" binding:"required"`
	LoginRequest
	Email string `json:"email" binding:"required,email"`
}
