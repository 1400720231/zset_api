package user

import (
	"fmt"
	"github.com/beego/beego/v2/server/web/context"
	"strconv"

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
	userid, _ := self.GetInt("id")
	fmt.Println("UserinfoControlleruserid", userid)
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

func (self *UserinfoController) Authorization(ctx *context.Context) (int, error) {
	userId, err := common.JwtAuthorization(ctx)
	fmt.Println("Authorization", userId)
	return userId, err

}

//权限钩子函数 默认通过
func (self *UserinfoController) CheckPermission(ctx *context.Context) bool {
	//fmt.Println("self.GetString(\"user_id\") ", self.GetInt("id"))
	user_id_string := self.Ctx.Input.Param(":user_id")
	id, _ := strconv.Atoi(user_id_string)
	userid := self.Ctx.Input.GetData("user_id")
	fmt.Println("CheckPermission", id, userid)
	if id != userid {
		return false
	}
	status := common.IsLoginPermission(ctx)
	fmt.Println(status)
	return status
}
