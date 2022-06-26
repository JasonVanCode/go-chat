package ws

import (
	"encoding/json"
)

//brodcat处理的消息
type Message struct {
	Sender    int64       `json:"sender,omitempty"`
	Recipient string      `json:"recipient,omitempty"`
	Content   string      `json:"content,omitempty"`
	Mes       interface{} `json:"mes"`
}

//聊天消息
type ChatMsg struct {
	Code        int    `json:"code,omitempty"`
	FromId      int    `json:"from_id,omitempty"`
	Msg         string `json:"msg,omitempty"`
	ToId        int    `json:"to_id,omitempty"`
	Status      int    `json:"status,omitempty"`
	MsgType     int    `json:"msg_type,omitempty"`
	ChannelType int    `json:"channel_type"`
}

func (msg *ChatMsg) ToJson() string {
	str, _ := json.Marshal(msg)
	return string(str)
}

//心跳返回的消息
type HeartBeatMsg struct {
	Code int    `json:"code"`
	Data string `json:"data"`
}

func NewHeartBeatMsg(code int, data string) *HeartBeatMsg {
	return &HeartBeatMsg{
		code,
		data,
	}
}

func (msg *HeartBeatMsg) ToJson() string {
	str, _ := json.Marshal(msg)
	return string(str)
}
