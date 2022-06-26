package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

//下面是协程的测试
func main222() {
	a := make(chan struct{})

	go func() {
		for {
			time.Sleep(time.Second)
			select {
			case <-a:
				fmt.Println("aaaa")
			}
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second)
			select {
			case a <- struct{}{}:
			}
		}
	}()

	time.Sleep(10 * time.Second)
}

//定义消息格式
type Msg struct {
	Code        int    `json:"code,omitempty"`
	FromId      int    `json:"from_id,omitempty"`
	Msg         string `json:"msg,omitempty"`
	ToId        int    `json:"to_id,omitempty"`
	Status      int    `json:"status,omitempty"`
	MsgType     int    `json:"msg_type,omitempty"`
	ChannelType int    `json:"channel_type"`
}

type ImClient struct {
	ID     int64           //客户端id
	Socket *websocket.Conn //
	Send   chan []byte
	Mux    sync.RWMutex
}

//brodcat处理的消息
type Message struct {
	Sender    int64  `json:"sender,omitempty"`
	Recipient string `json:"recipient,omitempty"`
	Content   string `json:"content,omitempty"`
	Mes       *Msg
}

func (c *ImClient) HandleMsg(msg []byte) {
	if len(msg) == 0 {
		return
	}
	//心跳，防止断开
	if string(msg) == "HeartBeat" {
		c.HandleMsgAck([]byte(`{"code":0,"data":"heartbeat ok"}`))
		return
	}

	var msgStruct Msg
	err := json.Unmarshal(msg, &msgStruct)
	if err != nil {
		fmt.Println("msg type is not right")
		return
	}

	if msgStruct.ChannelType == 1 {
		data := fmt.Sprintf(`{"code":200,"msg":"%s","from_id":%v,"to_id":%v,"status":"0","msg_type":%v,"channel_type":%v}`,
			msgStruct.Msg, msgStruct.FromId, msgStruct.ToId, msgStruct.MsgType, msgStruct.ChannelType)
		c.HandleMsgAck([]byte(data))
	}
	//消息广播
	ImManager.Broadcast <- &Message{Sender: c.ID, Mes: &msgStruct}

}

//当前
func (c *ImClient) HandleMsgAck(msg []byte) {
	c.Socket.WriteMessage(websocket.TextMessage, msg)
	return
}

func (c *ImClient) Read() {
	defer c.Socket.Close()
	for {
		_, msg, err := c.Socket.ReadMessage()
		if err != nil {
			fmt.Println("客户端连接失败，表示终端")
			ImManager.Unregister <- c
			c.Socket.Close()
			return
		}
		//处理消息
		c.HandleMsg(msg)
	}

}

func (c *ImClient) Write() {
	defer c.Socket.Close()
	for {
		select {
		case msg, ok := <-c.Send:
			if !ok {
				c.Socket.WriteMessage(websocket.CloseMessage, nil)
				c.Socket.Close()
				break
			}
			c.Socket.WriteMessage(websocket.TextMessage, msg)
		}
	}

}

type ImClientManager struct {
	ImClientMap map[int64]*ImClient // 存放在线用户连接
	Broadcast   chan *Message       // 收集消息分发到客户端
	Register    chan *ImClient      // 新注册的长连接
	Unregister  chan *ImClient      // 已注销的长连接
	lock        sync.Mutex
}

var ImManager = ImClientManager{
	ImClientMap: make(map[int64]*ImClient),
	Broadcast:   make(chan *Message),
	Register:    make(chan *ImClient),
	Unregister:  make(chan *ImClient),
}

//将客户端的信息保存
func (m *ImClientManager) SetCLientInfo(client *ImClient) {
	m.lock.Lock()
	m.ImClientMap[client.ID] = client
	m.lock.Unlock()
}

//消息处理
func (m *ImClientManager) HandleClientMsg(msg *Message) {
	//获取到需要发送到哪个客户端id
	to_id := msg.Mes.ToId
	//当前客户端在不在
	if client, ok := m.ImClientMap[int64(to_id)]; ok {
		json_msg, _ := json.Marshal(msg.Mes)
		client.Send <- json_msg
	}
}

//将字符串转int
func TransStringToInt(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return v
}

func main() {
	go func() {
		for {
			select {
			//要注册的客户端信息
			case client := <-ImManager.Register:
				fmt.Printf("处理了客户端%d的连接", client.ID)
				ImManager.SetCLientInfo(client) //将客户端信息保存
			case msg := <-ImManager.Broadcast:
				//负责消息传递
				ImManager.HandleClientMsg(msg)
			case client := <-ImManager.Unregister:
				//处理断开连接的客户端
				fmt.Println(client.ID, "连接断开了")

			}
		}

	}()

	r := gin.Default()
	r.GET("/ws", func(c *gin.Context) {
		id := TransStringToInt(c.Query("id"))
		fmt.Println("获取到当前id值是---->", id)
		if websocket.IsWebSocketUpgrade(c.Request) {
			conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
			if err != nil {
				c.Writer.Write([]byte("fail"))
				return
			}
			client := &ImClient{
				ID:     int64(id),
				Socket: conn,
				Send:   make(chan []byte),
			}
			//将要注册的协程推入管道
			ImManager.Register <- client
			//开启常驻协程发送数据
			go client.Read()
			go client.Write()
		}
	})
	r.Run(":9505") // listen and serve on 0.0.0.0:8080
}
