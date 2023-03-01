package request

type Register struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (register Register) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"Name.required":     "用户名不能为空",
		"Email.required":    "邮箱不能为空",
		"Password.required": "密码不能为空",
	}
}
