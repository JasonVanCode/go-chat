package ws

import (
	"fmt"
	"gin-admin/common/helper"
	"gin-admin/initialize"
	"github.com/streadway/amqp"
	"sync"
)

var WsManager *WsClientManager

//处理客户端的连接
type WsClientManager struct {
	ImClientMap map[int64]*WsClient // 存放在线用户连接
	Broadcast   chan *Message       // 收集消息分发到客户端
	Register    chan *WsClient      // 新注册的长连接
	Unregister  chan *WsClient      // 已注销的长连接
	lock        sync.Mutex
}

func NewWsClientManager() *WsClientManager {
	return &WsClientManager{
		ImClientMap: make(map[int64]*WsClient),
		Register:    make(chan *WsClient),
		Unregister:  make(chan *WsClient),
		Broadcast:   make(chan *Message),
	}
}

//将客户端信息保存map中
func (manager *WsClientManager) SaveClientMap(client *WsClient) {
	manager.lock.Lock()
	manager.ImClientMap[client.Id] = client
	manager.lock.Unlock()
}

//客户端离线，删除map中的信息
func (manager *WsClientManager) DelClientMap(client *WsClient) {
	//断开mq的连接
	manager.lock.Lock()
	delete(manager.ImClientMap, client.Id)
	manager.lock.Unlock()
}

//上线通知

//下线通知

//下发消息
func (manager *WsClientManager) LuanchMessage(msg *Message) {
	msgType := msg.Mes
	switch m := msgType.(type) {
	//聊天消息
	case *ChatMsg:
		manager.LuanchPrivateChatMsg(m)
	//心跳消息
	case *HeartBeatMsg:
		manager.LuanchAckMsg(msg.Sender, m)
	}
}

//返回自己客户端消息
func (manager *WsClientManager) LuanchAckMsg(client_id int64, msg *HeartBeatMsg) {
	if client, ok := manager.ImClientMap[client_id]; ok {
		fmt.Println("jajaj这是心跳消息")
		msgByte := []byte(msg.ToJson())
		client.Send <- msgByte
	}
}

//发送私聊消息
func (manager *WsClientManager) LuanchPrivateChatMsg(msg *ChatMsg) {
	//在线消息
	if client, ok := manager.ImClientMap[int64((msg.ToId))]; ok {
		msgByte := []byte(msg.ToJson())
		client.Send <- msgByte
	} else {
		fmt.Println("客户端", msg.ToId, "下线了，发送到rabbitnq中")
		//离线消息发送到rabiitmq的消息
		manager.SynPrivateMsgToChannel(msg)
	}
}

//同步私聊消息rabbitmq中
func (manager *WsClientManager) SynPrivateMsgToChannel(msg *ChatMsg) {
	//创建信道 channel 操作 队列
	ch, err := initialize.MqCoon.Channel()
	if err != nil {
		fmt.Println("发送channel失败", err)
		return
	}
	defer ch.Close()

	//声明要操作的队列
	queenName := "PersonalChat" + helper.TransIntToStirng(msg.ToId)
	q, err := ch.QueueDeclare(
		queenName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		fmt.Println("发送channel失败", err)
		return
	}
	//发送消息
	fmt.Println("往channel中发送了", msg.Msg, queenName)
	err = ch.Publish("", q.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(msg.Msg),
	})
	fmt.Println("发送channel失败", err)
}

//获取离线的私聊消息
func (manager *WsClientManager) SynOfflinePrivateMsgFromChannel(client *WsClient) {
	fmt.Println("处理同步消息")
	//创建信道 channel 操作 队列
	ch, err := initialize.MqCoon.Channel()
	if err != nil {
		fmt.Println("处理同步失败---》", err)
		return
	}

	//声明要操作的队列
	queenName := "PersonalChat" + helper.TransIntToStirng(int(client.Id))
	q, err := ch.QueueDeclare(
		queenName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		fmt.Println("处理同步失败---》", err)
		return
	}
	msg, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		fmt.Println("处理同步失败---》", err)
		return
	}

	//开启协程同步离线消息
	go func(ch *amqp.Channel) {
		defer ch.Close()
		for {
			select {
			case v, ok := <-msg:
				if ok {
					fmt.Println("获取到channel的离线消息是----------->", string(v.Body))
					client.Send <- v.Body
				} else {
					return
				}
			case <-client.StoMq: //客户端断开连接，停止监听mq
				fmt.Println("断开了mq的连接")
				return
			}
		}
	}(ch)
}

//发送群聊消息

//处理客户端的连接，消息的发送
func (manager *WsClientManager) Run() {
	for {
		select {
		case client := <-manager.Register:
			fmt.Println("客户端id为", client.Id, "连接进来了")
			//保存连接客户端的消息
			manager.SaveClientMap(client)
			//获取离线消息
			fmt.Println("协程池---》", initialize.AntsPool)
			initialize.AntsPool.Submit(func() {
				manager.SynOfflinePrivateMsgFromChannel(client)
			})
		case client := <-manager.Unregister:
			fmt.Println("客户端id为", client.Id, "连接断开")
			manager.DelClientMap(client)

		case msg := <-manager.Broadcast:
			fmt.Println("处理消息", msg)
			manager.LuanchMessage(msg)
		}
	}
}
