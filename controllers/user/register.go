package user

import (
	"fmt"
	"github.com/beego/beego/v2/adapter/orm"
	"zset_api/common"
	"zset_api/models/user"
)

//login的Controller
type RegisterController struct {
	common.BaseController
}

type Register struct {
	Username  string `json:"username"`  //前端传的字段名必须是username
	Password1 string `json:"password1"` //前端传的字段名必须是password
	Password2 string `json:"password2"` //前端传的字段名必须是password
}

//利用结构体的绑定实现http的post方法 beego内部对应的Post函数
func (self *RegisterController) Post() {

	username := self.GetString("username")
	password1 := self.GetString("password1")
	//计算密码的md5值
	hash_pwd, _ := common.HashAndSalt(password1)
	//保存数据库
	fmt.Println("hash_pwd", hash_pwd)

	user := &user.User{Username: username, Password: hash_pwd}
	//声明orm对象
	o := orm.NewOrm()
	o.Insert(user) //插入数据库
	results := map[string]interface{}{"code": 200, "message": "注册成功", "data": ""}
	self.Data["json"] = results

}

//路由级别权限校验
func (self *RegisterController) CheckPermission() bool {
	userid, _ := self.GetInt("userid")
	if self.UserId != userid {
		return false
	}
	return true

}

//form表单校验
func (self *RegisterController) CheckValidPermission() {
	username := self.GetString("username")
	password1 := self.GetString("password1")
	password2 := self.GetString("password2")
	results := map[string]interface{}{"code": 200} //默认状态码200
	if password1 != password2 {
		results = map[string]interface{}{"code": 600, "message": "两次密码错误", "data": ""}

	}
	//声明orm对象
	o := orm.NewOrm()
	ok := o.QueryTable("user").Filter("username", username).Exist()
	if ok == true {
		results = map[string]interface{}{"code": 600, "message": "用户名已经被使用", "data": ""}
	} else {
		results = map[string]interface{}{"code": 200, "message": "", "data": ""}

	}
	self.Data["json"] = results
	self.ServeJSON()

}
