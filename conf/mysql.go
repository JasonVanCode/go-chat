package conf

import "gin-admin/conf/viper"

func init() {
	viper.Add("mysql", viper.StrMap{
		"host":     viper.Env("DB_HOST", "127.0.0.1"),
		"port":     viper.Env("DB_PORT", "3306"),
		"database": viper.Env("DB_DATABASE", "im"),
		"user":     viper.Env("DB_USERNAME", "root"),
		"password": viper.Env("DB_PASSWORD", "root"),
		"charset":  viper.Env("DB_CHARSET", "utf8mb4"),
	})
}
