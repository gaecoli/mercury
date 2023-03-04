package bootstrap

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"mercury/mercury/utils"
	"reflect"
	"strings"
)

func InitializeValidator() {
	if vv, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = vv.RegisterValidation("email", utils.ValidateEmail)

		vv.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}
}
