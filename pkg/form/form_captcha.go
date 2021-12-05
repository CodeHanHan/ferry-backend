package form

type CaptchaResponse struct {
	Code string `json:"code"`
	Data string `json:"data"`
	Id   string `json:"id"`
	Msg  string `json:"msg"`
}

type VerifyCaptchaRequest struct {
	Id   string `json:"id" form:"id" binding:"required"`
	Code string `json:"code" form:"code" binding:"required"`
}
