package route

import (
	"gin-admin/core/http/controller"
	"gin-admin/core/http/middleware"
	"github.com/gin-gonic/gin"
)

//注册路由
func RegisterChatRoutes(engine *gin.Engine) {
	ws := new(controller.WsChatController)
	api := engine.Group("/im", middleware.Auth())
	{
		//聊天
		api.GET("/chat", ws.Connect)
	}

}
