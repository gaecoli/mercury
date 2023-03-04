package services

import (
	"errors"
	"mercury/mercury/app/common/request"
	"mercury/mercury/global"
	"mercury/mercury/models"
	"mercury/mercury/utils"
)

type userService struct {
}

var UserService = new(userService)

func (userService *userService) Register(params request.Register) (err error, user models.User) {
	var result = global.App.DB.Where("email = ?", params.Email).Select("id").First(&models.User{})
	if result.RowsAffected != 0 {
		err = errors.New("邮箱已存在")
		return
	}
	user = models.User{Name: params.Name, Email: params.Email, Password: utils.BcryptMake([]byte(params.Password))}
	err = global.App.DB.Create(&user).Error
	return
}
