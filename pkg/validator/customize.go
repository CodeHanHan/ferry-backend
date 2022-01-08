package validator

import (
	"github.com/go-playground/validator/v10"
)

var roleIsAdminValidator validator.Func = func(f1 validator.FieldLevel) bool {
	isAdmin, ok := f1.Field().Interface().(int)
	if !ok {
		return false
	}

	if isAdmin < 0 || isAdmin > 1 {
		return false
	}

	return true
}
