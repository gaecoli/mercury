package request

type Register struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (register Register) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"name.required":     "用户名不能为空",
		"email.required":    "邮箱不能为空",
		"email.email":       "邮箱格式不正确",
		"password.required": "密码不能为空",
	}
}
