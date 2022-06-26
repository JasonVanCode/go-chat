package controller

import "github.com/gin-gonic/gin"

type UserController struct {
}

func (u *UserController) Index(c *gin.Context) {

	c.JSON(200, struct {
		Msg string
	}{
		"hahaha",
	})
}
