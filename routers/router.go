package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"zset_api/controllers"
	"zset_api/controllers/user"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &user.LoginController{})
	//这里的int不是帮你强制转换成int 只是正则表达式的快捷写法，也就是说匹配出来的还是字符串的数字串
	beego.Router("/userinfo/:user_id:int", &user.UserinfoController{})
	beego.Router("/register", &user.RegisterController{})

}
