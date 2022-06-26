package middleware

import (
	"fmt"
	"gin-admin/common/jwt"
	"gin-admin/common/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

var (
	UserId     int //用户id
	AuthExcept = map[string]int{
		"/api/login": 1,
	}
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !IsAuthverify(c) {
			c.Next()
			return
		}
		//获取token
		token, ok := c.Request.Header["Authorization"]
		if !ok {
			response.ErrResponse(http.StatusUnauthorized, "验证失败", nil).ToJson(c)
			c.Abort()
			return
		}
		tokenStr := HandleBearToken(token[0])
		payload, err := jwt.ParseToken(tokenStr)
		if err != nil {
			response.ErrResponse(http.StatusUnauthorized, "验证失败", nil).ToJson(c)
			c.Abort()
			return
		}
		//获取到当前用户id
		UserId = payload.UserID
		fmt.Println("获取到当前用户的id是", UserId)
		c.Next()
	}

}

//是否登录验证
func IsAuthverify(c *gin.Context) bool {
	//当前请求路径
	urlSlipt := strings.Split(c.Request.RequestURI, "/")

	url := strings.Join(urlSlipt, "/")

	if _, ok := AuthExcept[url]; ok {
		return false
	}
	return true
}

//解析bear token
func HandleBearToken(token string) string {
	bearToken := strings.Split(token, " ")
	if len(bearToken) == 2 {
		return bearToken[1]
	}
	return ""
}
