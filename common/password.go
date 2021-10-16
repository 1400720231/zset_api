package common

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"golang.org/x/crypto/bcrypt"
)

//加密
func HashAndSalt(pwd string) (string, error) {
	byteHash := []byte(pwd)
	hash, err := bcrypt.GenerateFromPassword(byteHash, bcrypt.MinCost)
	if err != nil {
		fmt.Println("xx")
		return "", err
	}
	return string(hash), nil
}

//密码验证
func ComparePwd(hashPwd string, plainPwd []byte) bool {
	logs.Info(hashPwd)
	byteHash := []byte(hashPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		logs.Error(err.Error())
		return false
	}

	return true
}
