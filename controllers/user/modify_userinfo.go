package user

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web/context"
	"strconv"
	"zset_api/common"
	"zset_api/models/user"
)

// RegisterController login的Controller
type ModifyUserinfoController struct {
	common.BaseController
}

//form表单
type UserinfoForm struct {
	Age     int    `form:"age"`     //前端传的字段名必须是age
	Gender  int    `form:"gender"`  //前端传的字段名必须gender
	Email   string `form:"email"`   //前端传的字段名必须email
	Profile string `form:"profile"` //前端传的字段名必须profile
}

//application/json表单
type UserinfoJson struct {
	Age     int    `json:"age"`     //前端传的字段名必须是age
	Gender  int    `json:"gender"`  //前端传的字段名必须gender
	Email   string `json:"email"`   //前端传的字段名必须email
	Profile string `json:"profile"` //前端传的字段名必须profile
}

// Post 利用结构体的绑定实现http的post方法 beego内部对应的Post函数
func (self *ModifyUserinfoController) Post() {

	userIns := user.User{} //model 对象
	form := UserinfoForm{} //表单对象

	o := orm.NewOrm()
	o.QueryTable("user").Filter("id", self.UserId).One(&userIns)
	fmt.Println(userIns)
	//json_form =UserinfoJson{}
	//err := json.Unmarshal(self.Ctx.Input.RequestBody, &json_form)//json解析 也就是前端application/json
	//https://beego.me/docs/mvc/controller/params.md 这里要用取指针操作
	err := self.ParseForm(&form) //form表单解析成结构体
	fmt.Println(userIns)
	fmt.Println(err)
	if err == nil {

		userIns.Age = form.Age
		userIns.Gender = form.Gender
		userIns.Email = form.Email
		userIns.Profile = form.Profile
		fmt.Println(userIns)
		//更新数据库 确保主键在user.Id才能执行update
		//update更新&user的全部数据
		//o.Update(&user)//[UPDATE `user` SET `user_name` = ?, `password` = ?, `age` = ?, `gender` = ?, `phone` = ?, `email` = ?, `is_active` = ?, `is_delete` = ?, `profile` = ?, `update_time` = ? WHERE `id` = ?]
		//，可以选择更新部分数据o.Update(&user, "Field1", "Field2", ...)
		o.Update(&userIns, "Age", "Gender", "Email", "Profile") //[UPDATE `user` SET `age` = ?, `gender` = ?, `email` = ?, `profile` = ?, `update_time` = ? WHERE `id` = ?]
	}
	results := map[string]interface{}{"code": 200, "message": "修改数据成功", "data": ""}
	self.Data["json"] = results
	self.ServeJSON()

}

//认证
func (self *ModifyUserinfoController) Authorization(ctx *context.Context) (int, error) {
	userId, err := common.JwtAuthorization(ctx)
	fmt.Println("Authorization", userId)
	return userId, err

}

// 权限
func (self *ModifyUserinfoController) CheckPermission(ctx *context.Context) bool {
	user_id := self.Ctx.Input.Param(":user_id") //注意这里的从路由中解析参数需要加 “:”
	fmt.Println("ModifyUserinfoController", user_id)
	fmt.Println(self.UserId)
	UserId, _ := strconv.Atoi(user_id)
	if UserId != self.UserId {
		return false
	}
	return true
}

//业务校验
func (self *ModifyUserinfoController) CheckValidPermission(ctx *context.Context) map[string]interface{} {
	return map[string]interface{}{"code": 200}
}
