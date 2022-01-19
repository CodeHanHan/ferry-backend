package user

import (
	"github.com/CodeHanHan/ferry-backend/pkg/app"
	"github.com/CodeHanHan/ferry-backend/pkg/captcha"
	formUser "github.com/CodeHanHan/ferry-backend/pkg/form/user"
	"github.com/CodeHanHan/ferry-backend/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

var store = base64Captcha.DefaultMemStore

// Captcha godoc
// @Summary 获取验证码
// @Description 获取验证码
// @Tags captcha
// @ID get-captcha
// @Success 200 {object} formUser.CaptchaResponse
// @Failure 500 {object} app.ErrResponse
// @Produce  json
// @Router /captcha [get]
func Captcha(c *gin.Context) {
	id, b64s, err := captcha.DriverDigitFunc()
	if err != nil {
		logger.Error(c, "验证码生成失败")
		app.InternalServerError(c)
		return
	}

	app.Custom(c, gin.H{
		"code": 200,
		"data": b64s,
		"id":   id,
		"msg":  "success",
	})
}

// Captcha godoc
// @Summary 验证验证码
// @Description 验证验证码
// @Tags captcha
// @ID verify-captcha
// @Param id query string true "验证码id"
// @Param code query string true "验证码内容"
// @Success 200 {object} formUser.VerifyCaptchaResponse
// @Failure 400 {object} app.ErrResponse
// @Produce  json
// @Router /captcha [post]
func VerifyCaptcha(c *gin.Context) {
	var req formUser.VerifyCaptchaRequest
	if err := c.ShouldBind(&req); err != nil {
		logger.Error(c, "参数验证失败")
		app.ErrorParams(c, err)
		return
	}

	ok := store.Verify(req.Id, req.Code, true)
	if ok {
		app.OK(c, formUser.VerifyCaptchaResponse{
			Result: "success",
		})
		return
	} else {
		app.OK(c, formUser.VerifyCaptchaResponse{
			Result: "fail",
		})
		return
	}
}
