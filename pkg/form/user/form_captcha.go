package user

// get: /captcha
type CaptchaResponse struct {
	Code string `json:"code"`
	Data string `json:"data"`
	Id   string `json:"id"`
	Msg  string `json:"msg"`
}

// post /captcha
type VerifyCaptchaRequest struct {
	Id   string `json:"id" form:"id" binding:"required"`
	Code string `json:"code" form:"code" binding:"required"`
}

type VerifyCaptchaResponse struct {
	Result string `json:"result"`
}
