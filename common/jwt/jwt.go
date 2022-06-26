package jwt

import (
	"errors"
	"fmt"
	"gin-admin/conf/viper"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

// JWT : HEADER PAYLOAD SIGNATURE
var (
	SecretKEY              string = viper.GetString("jwt.secretkey")
	DEFAULT_EXPIRE_SECONDS int    = viper.GetInt("jwt.expireseconds") // default expired 3 minute
	PasswordHashBytes      int    = viper.GetInt("jwt.hashbytes")
)

type MyCustomClaims struct {
	UserID int `json:"UserID"`
	jwt.StandardClaims
}

// This struct is the parsing of token payload
type JwtPayload struct {
	Username  string `json:"Username"`
	UserID    int    `json:"UserID"`
	IssuedAt  int64  `json:"Iat"`
	ExpiresAt int64  `json:"Exp"`
}

//生成jwt token
func GenerateToken(userId int, userName string, expiredSeconds int) (string, error) {
	//token失效时间
	if expiredSeconds == 0 {
		expiredSeconds = DEFAULT_EXPIRE_SECONDS
	}

	secretKey := []byte(SecretKEY)
	issuedAt := time.Now().Unix()
	expiresAt := time.Now().Add(time.Second * time.Duration(expiredSeconds)).Unix()
	//声明claims
	claims := MyCustomClaims{
		userId,
		jwt.StandardClaims{
			Issuer:    userName,
			IssuedAt:  issuedAt,
			ExpiresAt: expiresAt,
		},
	}
	//Create the token using your claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//Signs the token with a secret
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", errors.New("generate token fail")
	}

	return tokenString, nil
}

//解密jwt token
func ParseToken(token string) (*JwtPayload, error) {
	tokenObj, _ := jwt.ParseWithClaims(token, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKEY), nil
	})
	claims, ok := tokenObj.Claims.(*MyCustomClaims)

	if ok && tokenObj.Valid {
		fmt.Println("用户id", claims.UserID)
		fmt.Println("开始时间", claims.IssuedAt)
		fmt.Println("结束时间", claims.ExpiresAt)
		mowTime := time.Now().Unix()
		fmt.Println("当前时间：", mowTime)
		return &JwtPayload{
			claims.Issuer,
			claims.UserID,
			claims.IssuedAt,
			claims.ExpiresAt,
		}, nil
	}
	return nil, errors.New("Illegal token")
}

//刷新token
func RefreshToken(token string) (string, error) {
	tokenObj, _ := jwt.ParseWithClaims(token, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKEY), nil
	})

	claims, ok := tokenObj.Claims.(*MyCustomClaims)
	if !ok || !tokenObj.Valid {
		return "", errors.New("token expiration or Illegal token")
	}
	return GenerateToken(claims.UserID, claims.Issuer, 0)
}
