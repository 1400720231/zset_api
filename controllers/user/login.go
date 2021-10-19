package user

import (
	"fmt"
	"github.com/beego/beego/v2/adapter/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web/context"
	"zset_api/common"
	"zset_api/models/user"
)

//login的Controller
type LoginController struct {
	common.BaseController
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//利用结构体的绑定实现http的post方法 beego内部对应的Post函数
func (self *LoginController) Post() {
	data := make(map[string]interface{})
	response := map[string]interface{}{"code": 200, "message": "", "data": ""} //默认状态码200
	username := self.GetString("username")                                     //默认值为空字符串
	password := self.GetString("password")
	fmt.Println(username, password) //默认值为空字符串
	md5Pwd := common.GetMd5Str(password)
	userStruct := &user.User{}
	//声明orm对象
	o := orm.NewOrm()
	//查询数据库是否有当前用户和密码对应
	err := o.QueryTable("user").Filter("user_name", username).Filter("password", md5Pwd).One(userStruct)
	logs.Info("LoginController", err)
	if err == nil {
		token, _ := common.GenerateJwtToken(username, userStruct.Id)
		data["user_id"] = userStruct.Id
		data["username"] = userStruct.Username
		data["token"] = token
		response["data"] = data
	}
	self.Data["json"] = response
	self.ServeJSON()

}

// 不做认证
func (self *LoginController) Authorization(ctx *context.Context) (int, error) {
	return 0, nil
}

// 认通过
func (self *LoginController) CheckPermission(ctx *context.Context) bool {
	return true
}

// CheckValidPermission form表单校验
func (self *LoginController) CheckValidPermission(ctx *context.Context) map[string]interface{} {
	response := map[string]interface{}{"code": 200, "message": "", "data": ""} //默认状态码200
	username := self.GetString("username")
	password := self.GetString("password")
	md5Pwd := common.GetMd5Str(password)
	//声明orm对象
	o := orm.NewOrm()
	ok := o.QueryTable("user").Filter("user_name", username).Filter("password", md5Pwd).Exist()
	logs.Info("LoginController", ok)
	if ok {

		return response
	}

	response["code"] = 600
	response["message"] = "账号或密码错误"

	return response
}

/*
//fmt.Printlnet的request")
	//fmt.Println(self.Ctx.Request.Proto)      //HTTP/1.1
	//fmt.Println(self.Ctx.Request.URL)        ///login?page=1
	//fmt.Println(self.Ctx.Request.RequestURI) ///login?page=1
	//fmt.Println(self.Ctx.Request.Header)     //map类型的headers
	//fmt.Println(self.Ctx.Request.RemoteAddr) //127.0.0.1:58565
	//fmt.Println(self.Ctx.Request.Host)       //127.0.0.1:8080 （beego项目的启动域名 线上应该是域名）
	//fmt.Println(self.Ctx.Request.Method)     //POST
	fmt.Println(self.Ctx.Request.Form)       //map[page:[1] username:[panda]]
	fmt.Println(self.Ctx.Request.PostForm)   //map[username:[panda]]
	//fmt.Println(self.Ctx.Request.ContentLength)//68
	//fmt.Println(self.Ctx.Request.URL.Path)///login
	//fmt.Println(self.Ctx.Request.URL.RawQuery)//page=1

	fmt.Println(self.Ctx.Input.Refer())
	fmt.Println(self.Ctx.Input.Port())
	fmt.Println(self.Ctx.Input.SubDomains())
	fmt.Println(self.Ctx.Input.UserAgent())
	fmt.Println("ParamsLen", self.Ctx.Input.ParamsLen())
	fmt.Println("Param", self.Ctx.Input.Param("page"))
	fmt.Println("Params", self.Ctx.Input.Params())
	fmt.Println("Cookie", self.Ctx.Input.Cookie("page"))
	//fmt.Println(self.Ctx.Input.Session("page"))
	fmt.Println("Data", self.Ctx.Input.Data())
	fmt.Println("GetData", self.Ctx.Input.GetData("username"))
	fmt.Println("GetData", self.Ctx.Input.GetData("page"))
	fmt.Println(self.Ctx.Input.Header("Content-Type"))
	fmt.Println(self.Ctx.Input.Cookie("page"))
	username:=self.Ctx.Request.PostForm["username"][0]
	fmt.Println(reflect.TypeOf(username))

	var newValues url.Values
	if self.Ctx.Request.URL != nil {

		newValues, _ = url.ParseQuery(self.Ctx.Request.URL.RawQuery)
		fmt.Println(newValues)}
	page:=newValues["page"]

	fmt.Println("pagepagepage,page",page[0])
	fmt.Println("net的request")
	fmt.Println("beego的request")
	//这里的Ctx是beego层面封装的不是net/http原生的
	//a,_:=self.Input()
	fmt.Println(self.Ctx.Input.Query("page"))
	fmt.Println(self.Ctx.Request.Form)
	fmt.Println(self.Ctx.Input.RequestBody)
	fmt.Println("beego的request")
*/
