package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:         "127.0.0.1:6379", //redis 连接地址
		Password:     "",               //密码
		DB:           0,
		ReadTimeout:  2 * time.Second, //socket 读超时时间
		WriteTimeout: 2 * time.Second, //socket 写超时时间
		DialTimeout:  2 * time.Second, //连接超时时间
		PoolSize:     15,              //连接池 默认为4倍cpu数
	})
	defer client.Close()

	go func() {
		psub := client.Subscribe("chan")
		for v := range psub.Channel() {
			fmt.Println(v.String())
		}
	}()

	timeer := time.NewTimer(5 * time.Second)
	for {
		time.Sleep(time.Second)
		select {
		case <-timeer.C:
			return
		default:
			client.Publish("chan", "我在发布消息")
		}
	}

}
