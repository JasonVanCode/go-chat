package conf

import (
	"fmt"
	"gin-admin/conf/viper"
)

func init() {
	fmt.Println("初始化rabbitmq的配置文件")
	viper.Add("rabbitmq", viper.StrMap{
		"host":     viper.Env("RABBITMQ_HOST", "127.0.0.1"),
		"port":     viper.Env("RABBITMQ_PORT", "5672"),
		"user":     viper.Env("RABBITMQ_USER", "guest"),
		"password": viper.Env("RABBITMQ_PASSWORD", "guest"),
	})
}
