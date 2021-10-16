package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"zset_api/controllers"
	"zset_api/controllers/user"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &user.LoginController{})
	beego.Router("/userinfo/:userid:int", &user.UserinfoController{})

}
