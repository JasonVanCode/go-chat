package mysql

import (
	"gin-admin/conf/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"time"
)

func InitMysql() *gorm.DB {
	var (
		host     = viper.GetString("mysql.host")
		port     = viper.GetString("mysql.port")
		database = viper.GetString("mysql.database")
		charset  = viper.GetString("mysql.charset")
		user     = viper.GetString("mysql.user")
		password = viper.GetString("mysql.password")
		err      error
	)

	dsn := strings.Join([]string{user, ":", password, "@tcp(", host, ":", port, ")/", database, "?charset=", charset, "&parseTime=True&loc=Local"}, "")
	DB, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{})
	if err != nil {
		panic("数据库连接异常")
	}
	sqlDB, err := DB.DB()
	if err != nil {
		panic("数据库连接异常")
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	return DB
}
