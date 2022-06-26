package controller

import (
	"gin-admin/common/ws"
	"gin-admin/core/http/middleware"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

type WsChatController struct {
	upgrader websocket.Upgrader
}

//处理websocket的连接
func (wsCon *WsChatController) Connect(c *gin.Context) {
	//判断是否是websocket
	if websocket.IsWebSocketUpgrade(c.Request) {
		con, err := ws.NewWebsocket(c)
		if err != nil {
			http.NotFound(c.Writer, c.Request)
			return
		}
		//当前用户的id
		userId := middleware.UserId
		client := ws.NewWsClient(int64(userId), con)
		//开启协程处理消息的读取和发送
		go client.WsRead()
		go client.WsWrite()
		//注册客户端信息
		ws.WsManager.Register <- client
	}

}
