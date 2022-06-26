package services

import (
	"errors"
	"gin-admin/common/helper"
	"gin-admin/common/jwt"
	"gin-admin/core/http/models"
	"gin-admin/core/http/validates"
	"gin-admin/initialize"
)

type UserService struct{}

//处理用户登录
func (*UserService) HandleLogin(body validates.UserLoginBody) (*models.User, string, error) {
	var DB = initialize.DB
	var user models.User
	DB.Where("name = ?", body.UserName).First(&user)
	if user.ID == 0 {
		return nil, "", errors.New("该用户不存在")
	}
	//验证密码
	isOk := helper.PasswordVerify(helper.TransStringToSliceByte(user.Password), helper.TransStringToSliceByte(body.PassWord))
	if isOk == false {
		return nil, "", errors.New("密码错误")
	}
	//生成jwt-token
	token, err := jwt.GenerateToken(user.ID, user.Name, 0)
	if err != nil {
		return nil, "", errors.New("token 验证失败")
	}
	return &user, token, nil
}
