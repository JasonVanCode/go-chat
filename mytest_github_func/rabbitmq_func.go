package main

import (
	"fmt"
	"github.com/streadway/amqp"
)

func main() {
	getMsg()
}

//rabbitmq发布消息
func publishMsg() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	//开启信道进行连接
	cha, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer cha.Close()
	body := "hahahaha "
	q, err := cha.QueueDeclare("test1", false, false, false, false, nil)
	err = cha.Publish("", q.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(body),
	})
	if err != nil {
		panic(err)
	}

}

//rabbitmq获取消息
func getMsg() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	//开启信道进行连接
	cha, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer cha.Close()
	q, err := cha.QueueDeclare("test1", false, false, false, false, nil)

	msg, err := cha.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	for v := range msg {
		fmt.Println(string(v.Body))
	}

	fmt.Println("hahaha")
	return
}
