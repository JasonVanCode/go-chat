package validates

import (
	"github.com/go-playground/validator/v10"
)

type UserLoginBody struct {
	UserName string `json:"user_name" validate:"required"`
	PassWord string `json:"pass_word" validate:"required"`
}

//验证请求
func (body *UserLoginBody) CheckValidate() error {
	validate := validator.New()
	return validate.Struct(body)
}
