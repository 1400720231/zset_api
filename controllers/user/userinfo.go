package user

import (
	"fmt"
	"zset_api/common"
)

type UserinfoController struct {
	common.BaseController
}
type UserInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//利用结构体的绑定实现http的post方法 beego内部对应的Post函数
func (self *UserinfoController) Get() {
	userid, _ := self.GetInt("userid")
	fmt.Println(userid)
	data := make(map[string]interface{})
	//序列化的数据库对象
	results := make(map[string]interface{})

	results["username"] = "boy"
	results["age"] = 18
	results["weight"] = 120
	//数据数组
	data["code"] = 200
	data["message"] = "success"
	data["results"] = results
	self.Data["json"] = data
	self.ServeJSON()
}
func (self *UserinfoController) CheckPermission() bool {
	userid, _ := self.GetInt("userid")
	if self.UserId != userid {
		return false
	}
	return true

}
