package rabbitmq

import (
	"fmt"
	"gin-admin/conf/viper"

	"github.com/streadway/amqp"
	"strings"
)

func InitMq() *amqp.Connection {
	var (
		user = viper.GetString("rabbitmq.user")
		pass = viper.GetString("rabbitmq.password")
		host = viper.GetString("rabbitmq.host")
		port = viper.GetString("rabbitmq.port")
	)
	dialUrl := strings.Join([]string{"amqp://", user, ":", pass, "@", host, ":", port}, "")
	fmt.Println("rabbitmq连接地址----->", dialUrl)
	MqCoon, err := amqp.Dial(dialUrl)
	if err != nil {
		panic("rabbitmq 连接失败")
	}
	return MqCoon
}
