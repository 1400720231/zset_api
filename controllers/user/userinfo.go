package user

import (
	"fmt"
	"github.com/beego/beego/v2/adapter/orm"
	"github.com/beego/beego/v2/server/web/context"
	"strconv"
	"zset_api/common"
	"zset_api/models/user"
)

type UserinfoController struct {
	common.BaseController
}
type UserInfo struct {
	Id string `json:"id"`
	//Username   string    `json:"username"`
	//Age        int       `json:"age"`
	//Gender     int       `json:"gender"`
	//Phone      int64     `json:"phone"`
	//Email      string    `json:"email"`
	//Profile    string    `json:"profile"`
	//CreateTime time.Time `json:"create_time"`
}

//利用结构体的绑定实现http的post方法 beego内部对应的Post函数
func (self *UserinfoController) Get() {
	data := make(map[string]interface{})
	response := map[string]interface{}{"code": 200, "message": "", "data": ""} //默认状态码200

	user_id := self.Ctx.Input.Param(":user_id")
	UserId, _ := strconv.Atoi(user_id)
	fmt.Println("UserinfoControlleruserid", UserId)

	userStruct := &user.User{}
	o := orm.NewOrm()

	//查询数据库是否有当前用户和密码对应 只返回部分字段值，其他全为默认值
	//select id,username from 而不是select * from
	o.QueryTable("user").Filter("id", UserId).Filter("is_delete", 0).One(userStruct, "Id", "Username", "age", "Gender", "Phone", "Email", "Profile", "CreateTime")
	Userinfo := userStruct.Userinfo()
	//数据数组
	data["code"] = 200
	data["message"] = "success"
	response["data"] = Userinfo
	self.Data["json"] = response
	self.ServeJSON()
}

//认证
func (self *UserinfoController) Authorization(ctx *context.Context) (int, error) {
	userId, err := common.JwtAuthorization(ctx)
	fmt.Println("Authorization", userId)
	return userId, err

}

// 权限
func (self *UserinfoController) CheckPermission(ctx *context.Context) bool {
	user_id := self.Ctx.Input.Param(":user_id") //注意这里的从路由中解析参数需要加 “:”
	fmt.Println("UserinfoController", user_id)
	fmt.Println(self.UserId)
	UserId, _ := strconv.Atoi(user_id)
	if UserId != self.UserId {
		return false
	}
	return true
}

//业务校验
func (self *UserinfoController) CheckValidPermission(ctx *context.Context) map[string]interface{} {
	return map[string]interface{}{"code": 200}
}
