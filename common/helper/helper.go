package helper

import (
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

//bcrypt 加密
func PasswordHash(password []byte) string {
	pass, _ := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	return TransSliceByteToString(pass)
}

//bcrypt 解密
func PasswordVerify(hashedPassword, password []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)
	return err == nil
}

//[]byte => 转 string
func TransSliceByteToString(param []byte) string {
	return string(param)
}

//string 转 []byte
func TransStringToSliceByte(param string) []byte {
	return []byte(param)
}

//int 转 字符串
func TransIntToStirng(param int) string {
	return strconv.Itoa(param)
}
