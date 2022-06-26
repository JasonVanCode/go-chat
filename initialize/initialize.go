package initialize

import (
	"gin-admin/initialize/antspool"
	"gin-admin/initialize/mysql"
	"gin-admin/initialize/rabbitmq"
	"github.com/panjf2000/ants/v2"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
)

//协程池连接
var AntsPool *ants.Pool

//数据库连接
var DB *gorm.DB

//rabbitmq 连接
var MqCoon *amqp.Connection

//初始化其他连接
func InitService() {
	AntsPool = antspool.InitPool()
	DB = mysql.InitMysql()
	MqCoon = rabbitmq.InitMq()
}
