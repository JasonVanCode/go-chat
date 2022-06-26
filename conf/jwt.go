package conf

import "gin-admin/conf/viper"

func init() {
	viper.Add("jwt", viper.StrMap{
		"secretkey":     viper.Env("JWT_SecretKEY", "JWT-Secret-Key"),
		"expireseconds": viper.Env("JWT_DEFAULT_EXPIRE_SECONDS", 180),
		"hashbytes":     viper.Env("JWT_PasswordHashBytes", 16),
	})
}
