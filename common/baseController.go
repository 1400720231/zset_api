package common

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"strings"
)

//https://beego.me/docs/mvc/controller/controller.md#%E5%AD%90%E7%B1%BB%E6%89%A9%E5%B1%95
type BaseControllerInterface interface {
	CheckPermission() bool                        //路由层面权限拦截
	CheckGetPermission() bool                     //get方法级别权限
	CheckPostPermission() bool                    //post方法级别权限
	CheckPutPermission() bool                     //put方法级别权限
	CheckDeletePermission() bool                  //delete方法级别权限
	CheckHeadPermission() bool                    //head方法级别权限
	CheckOptionsPermission() bool                 //options方法级别权限
	CheckPatchPermission() bool                   //Patch方法级别权限
	CheckValidPermission() map[string]interface{} //业务逻辑层面权限校验 比如验证码过期 您还没有权限操作
}

//login的Controller
type BaseController struct {
	beego.Controller
	UserId int //用户id 当然实际上不一定是int，具体看业务要求，一般不做分表分库单表的id肯定够

}

//重写beego中的Prepare钩子函数
func (self *BaseController) Prepare() {

	//auth := self.AppController.
	//从请求头中解析出token的参数
	authorization := self.Ctx.Input.Header("Authorization")
	//去空格
	token := strings.TrimSpace(authorization)
	//jwt 解析token返回解析体
	user, err := ParseToken(token)
	//如果err有错则说明解析有问题 需要重新登陆
	if err != nil {
		fmt.Println(err)
		self.Abort("401") //token校验失败
		//self.StopRun()
	} else { //解析没有错的话把当前request.user上
		self.UserId = user.UserId
	}

	//若果当前实例化的Controller是BaseControllerInterface
	//注意这里的self.AppController调用方法，必须app.CheckPermission()这样调用子类的方法，如果self.CheckPermission()
	//是没法找到子类CheckPermission的，这是因为beego的声明周期决定的
	//这样封装后只要子类可以会根据自己的业务逻辑重载对应的业务方法，注意CheckPermission返回的是bool
	//CheckValidPermission返回的是map[string]interface{}
	if app, ok := self.AppController.(BaseControllerInterface); ok {
		fmt.Println("CheckPermission")
		//权限校验 不通过直接403
		routerStatus := app.CheckPermission()
		if routerStatus == false {
			self.Abort("403")

		}
		getStatus := app.CheckGetPermission()
		if getStatus == false {
			self.Abort("403")

		}
		postStatus := app.CheckPostPermission()
		if postStatus == false {
			self.Abort("403")

		}
		putStatus := app.CheckPutPermission()
		if putStatus == false {
			self.Abort("403")

		}
		deleteStatus := app.CheckDeletePermission()
		if deleteStatus == false {
			self.Abort("403")

		}
		headStatus := app.CheckHeadPermission()
		if headStatus == false {
			self.Abort("403")

		}
		optionStatus := app.CheckOptionsPermission()
		if optionStatus == false {
			self.Abort("403")

		}

		validData := app.CheckValidPermission()
		fmt.Println("CheckValidPermission")
		//业务规则校验，默认校验通过
		if validData["code"] == 200 {
			return
		} else {
			//否则返回校验信息
			self.Data["json"] = validData
			self.ServeJSON()
		}
	}

	return
}

//封装权限
func (self *BaseController) CheckPermission() bool {
	return true
}

//get方法级别权限
func (self *BaseController) CheckGetPermission() bool {
	return true
}

//post方法级别权限
func (self *BaseController) CheckPostPermission() bool {
	return true
}

//put方法级别权限
func (self *BaseController) CheckPutPermission() bool {
	return true
}

//delete方法级别权限
func (self *BaseController) CheckDeletePermission() bool {
	return true
}

//head方法级别权限
func (self *BaseController) CheckHeadPermission() bool {
	return true
}

//options方法级别权限
func (self *BaseController) CheckOptionsPermission() bool {
	return true
}

//业务逻辑层面权限校验 比如验证码过期 您还没有权限操作
func (self *BaseController) CheckValidPermission() map[string]interface{} {
	return map[string]interface{}{"code": 200}
}
