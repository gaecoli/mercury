package utils

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

func ValidateEmail(vf validator.FieldLevel) bool {
	email := vf.Field().String()
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9.!#$%&'*+/=?^_` + "`" + `{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$`)

	// 排除 QQ 邮箱
	qqRegex := regexp.MustCompile(`@qq\.(com|com\.cn|cn)$`)
	if qqRegex.MatchString(email) {
		return false
	}

	// 验证邮箱格式是否符合规则
	return emailRegex.MatchString(email)
}
