package validator

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    ut.Translator
)

func SetUp() error {
	//注册翻译器
	zh := zh.New()
	uni = ut.New(zh, zh)

	trans, _ = uni.GetTranslator("zh")

	//获取gin的校验器
	validate := binding.Validator.Engine().(*validator.Validate)

	// 注册自定义校验器
	RegisterCustomizeValidator(validate)

	//注册翻译器
	return zh_translations.RegisterDefaultTranslations(validate, trans)
}

//Translate 翻译错误信息
func Translate(err error) interface{} {

	var result = make(map[string][]string)

	errors, ok := err.(validator.ValidationErrors)
	if !ok {
		return err
	}

	for _, err := range errors {
		result[err.Field()] = append(result[err.Field()], err.Translate(trans))
	}
	return result
}

// 注册自定义验证器
func RegisterCustomizeValidator(v *validator.Validate) {
	_ = v.RegisterValidation("is_admin", roleIsAdminValidator) // FIXME 错误处理
}
