package ws

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"sync"
)

type WsClient struct {
	Id    int64           //客户端id
	Conn  *websocket.Conn //websocket 客户端连接
	Send  chan []byte     //接收要发送的消息
	Lock  sync.Mutex
	StoMq chan struct{} // 如果客户端断开连接，停止从channel获取数据
}

//初始化客户端信息
func NewWsClient(id int64, con *websocket.Conn) *WsClient {
	return &WsClient{
		Id:    id,
		Conn:  con,
		Send:  make(chan []byte),
		StoMq: make(chan struct{}),
	}
}

//处理客户端消息
func (c *WsClient) HandleMsg(msg []byte) {
	//空消息不处理
	if len(msg) == 0 {
		return
	}
	var brodcatMsg Message
	brodcatMsg.Sender = c.Id
	//心跳处理
	if string(msg) == "heartBeat" {
		hearBeatMsg := NewHeartBeatMsg(200, "heartbeat ok")
		brodcatMsg.Mes = hearBeatMsg
		WsManager.Broadcast <- &brodcatMsg
		return
	}
	//聊天消息
	var chatMsg ChatMsg
	err := json.Unmarshal(msg, &chatMsg)
	if err != nil {
		fmt.Println("消息解析失败的原因----->", err)
		return
	}
	brodcatMsg.Mes = &chatMsg
	fmt.Println("开始往Broadcast发送消息")
	WsManager.Broadcast <- &brodcatMsg
}

//协程 处理websocket 读取消息
func (c *WsClient) WsRead() {
	defer func() {
		//处理连接中断的客户端信息
		WsManager.Unregister <- c
		c.StoMq <- struct{}{}
		_ = c.Conn.Close()
	}()
	for {
		//从客户端读取消息
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			fmt.Println("从客户端client", c.Id, "断开了连接")
			return
		}
		fmt.Println("获取客户端的数据hahhahaha->", string(msg))
		c.HandleMsg(msg)
	}
}

//协程 处理websocket 写消息
func (c *WsClient) WsWrite() {
	defer c.Conn.Close()
	for {
		select {
		//如果send chanel 获取到消息，需要给此客户端发送消息
		case msg, ok := <-c.Send:
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, nil)
				return
			}
			c.Conn.WriteMessage(websocket.TextMessage, msg)
		}
	}
}
