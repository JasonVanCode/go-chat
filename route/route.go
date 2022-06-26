package route

import (
	"gin-admin/core/http/controller"
	"gin-admin/core/http/middleware"
	"github.com/gin-gonic/gin"
)

//注册路由
func RegisterRoutes(engine *gin.Engine) {
	authController := new(controller.AuthController)
	userController := new(controller.UserController)
	api := engine.Group("/api", middleware.Auth())
	{
		//登录
		api.POST("/login", authController.Login)

		//用户查询
		userApi := api.Group("/user")
		{
			userApi.GET("index", userController.Index)
		}
	}

}
