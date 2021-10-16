package user

import (
	"zset_api/common"
)

//login的Controller
type LoginController struct {
	common.BaseController
}

type User struct {
	Username string `json:"username2"`
	Password string `json:"password"`
}

//利用结构体的绑定实现http的post方法 beego内部对应的Post函数
func (self *LoginController) Post() {
	data := make(map[string]interface{})
	//序列化的数据库对象
	results := make(map[string]interface{})
	//假设已经登陆了 当前用户的用户名为xiongyao userid=431
	token, _ := common.GenerateJwtToken("xiongyao", 431)
	results["token"] = token
	//数据数组
	data["code"] = 200
	data["message"] = "success"
	data["results"] = results
	self.Data["json"] = data
	self.ServeJSON()

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
