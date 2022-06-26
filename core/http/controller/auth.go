package controller

import (
	"gin-admin/common/response"
	"gin-admin/core/http/models"
	"gin-admin/core/http/validates"
	"gin-admin/core/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

//权限
type AuthController struct{}

func (auth *AuthController) Login(c *gin.Context) {
	var body validates.UserLoginBody
	if err := c.ShouldBindJSON(&body); err != nil {
		response.ErrResponse(http.StatusInternalServerError, err.Error(), nil).ToJson(c)
		return
	}
	if err := body.CheckValidate(); err != nil {
		response.ErrResponse(http.StatusInternalServerError, "请求参数有误", nil).ToJson(c)
		return
	}
	var userService services.UserService
	userInfo, token, err := userService.HandleLogin(body)
	if err != nil {
		response.ErrResponse(http.StatusInternalServerError, err.Error(), nil).ToJson(c)
		return
	}
	response.SuccessResponse(struct {
		User  *models.User
		Token string
	}{
		userInfo,
		token,
	}).ToJson(c)

}
