package services

import (
	"errors"
	"strconv"

	"Raven-Admin/app/common/request"
	"Raven-Admin/app/models"
	g "Raven-Admin/global"
	"Raven-Admin/utils"
)

type userService struct{}

var UserService = new(userService)

func (s *userService) Register(params request.Register) (err error, user models.User) {
	var result = g.DB.Where("username=?", params.Username).Select("id").First(&models.User{})

	if result.RowsAffected != 0 {
		err = errors.New("用户名已存在")
		return
	}

	user = models.User{Username: params.Username, Password: utils.BcryptMake([]byte(params.Password))}
	err = g.DB.Create(&user).Error
	return
}

func (s *userService) RegisterWithMobile(params request.RegisterWithMobile) (err error, user models.User) {
	var result = g.DB.Where("mobile=?", params.Mobile).Select("id").First(&models.User{})

	if result.RowsAffected != 0 {
		err = errors.New("手机号已存在")
		return
	}

	result = g.DB.Where("username=?", params.Username).Select("id").First(&models.User{})

	if result.RowsAffected != 0 {
		err = errors.New("用户名已存在")
		return
	}

	user = models.User{Username: params.Username, Mobile: params.Mobile, Password: utils.BcryptMake([]byte(params.Password))}
	err = g.DB.Create(&user).Error
	return
}

func (s *userService) Login(params request.Login) (err error, user models.User) {
	err = g.DB.Where("username=?", params.Username).First(&user).Error

	if err != nil || !utils.BcryptCheck(user.Password, []byte(params.Password)) {
		err = errors.New("账号密码错误")
	}

	return
}

func (s *userService) GetUserInfo(id string) (err error, user models.User) {
	intId, err := strconv.Atoi(id)
	err = g.DB.First(&user, intId).Error

	if err != nil {
		err = errors.New("用户不存在")
	}
	return
}

func (s *userService) GetUserInfoWhithGuardName(GuardName string, id string) (err error, user utils.JwtUser) {
	switch GuardName {
	case g.Cof.App.AppName:
		return s.GetUserInfo(id)
	default:
		err = errors.New("guard" + GuardName + "does not exist")

	}
	return
}
