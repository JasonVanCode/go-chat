package main

import (
	"fmt"
	"gin-admin/common/ws"
	"gin-admin/conf"
	"gin-admin/initialize"
	"gin-admin/route"
	"github.com/gin-gonic/gin"
)

func init() {
	//初始化配置文件
	conf.InitConfig()
}

func main() {
	//初始化mysql、rabbitmq、antspool
	initialize.InitService()
	////协程处理websocket 的连接处理
	ws.WsManager = ws.NewWsClientManager()
	go ws.WsManager.Run()
	engine := gin.Default()
	////注册路由
	route.RegisterRoutes(engine)
	////websocket路由
	route.RegisterChatRoutes(engine)
	fmt.Println("进程的池----->", initialize.AntsPool)

	engine.Run(":9095")
}
