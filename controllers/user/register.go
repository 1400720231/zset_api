package user

import (
	"fmt"
	"github.com/beego/beego/v2/adapter/orm"
	"github.com/beego/beego/v2/server/web/context"
	"zset_api/common"
	"zset_api/models/user"
)

// RegisterController login的Controller
type RegisterController struct {
	common.BaseController
}

type Register struct {
	Username  string `json:"username"`  //前端传的字段名必须是username
	Password1 string `json:"password1"` //前端传的字段名必须是password
	Password2 string `json:"password2"` //前端传的字段名必须是password
}

// Post 利用结构体的绑定实现http的post方法 beego内部对应的Post函数
func (self *RegisterController) Post() {

	username := self.GetString("username") //取不到默认空字符串
	password1 := self.GetString("password1")
	password2 := self.GetString("password2")
	fmt.Println(username)
	fmt.Println(password1)
	fmt.Println(password2)
	//计算密码的md5值
	md5Pwd := common.GetMd5Str(password1)
	//保存数据库
	fmt.Println("md5Pwd", md5Pwd)

	userStruct := &user.User{Username: username, Password: md5Pwd}
	//声明orm对象
	o := orm.NewOrm()
	_, err := o.Insert(userStruct) //插入数据库
	fmt.Println(err)
	results := map[string]interface{}{"code": 200, "message": "注册成功", "data": ""}
	self.Data["json"] = results
	self.ServeJSON()

}

// 不做认证
func (self *RegisterController) Authorization(ctx *context.Context) (int, error) {
	return 0, nil
}

// 认通过
func (self *RegisterController) CheckPermission(ctx *context.Context) bool {
	return true
}

// CheckValidPermission form表单校验
func (self *RegisterController) CheckValidPermission(ctx *context.Context) map[string]interface{} {
	username := self.GetString("username")
	password1 := self.GetString("password1")
	password2 := self.GetString("password2")

	results := map[string]interface{}{"code": 200, "message": "", "data": ""} //默认状态码200
	if password1 != password2 {
		results = map[string]interface{}{"code": 600, "message": "两次密码不一致", "data": ""}
		return results
	}
	//声明orm对象
	o := orm.NewOrm()
	ok := o.QueryTable("user").Filter("user_name", username).Exist()
	if ok == true {
		results = map[string]interface{}{"code": 600, "message": "用户名已经被使用", "data": ""}
	} else {
		results = map[string]interface{}{"code": 200, "message": "", "data": ""}

	}
	return results
}
