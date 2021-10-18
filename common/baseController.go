package common

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

//https://beego.me/docs/mvc/controller/controller.md#%E5%AD%90%E7%B1%BB%E6%89%A9%E5%B1%95
type BaseControllerInterface interface {
	//注意子类的函数名 和返回类型都要必须一模一样才叫做满足这个interface协议
	Authorization(ctx *context.Context) (int, error)                  //认证
	CheckPermission(ctx *context.Context) bool                        //路由层面权限拦截
	CheckValidPermission(ctx *context.Context) map[string]interface{} //业务逻辑层面权限校验 比如验证码过期 您还没有权限操作
}

//login的Controller
type BaseController struct {
	beego.Controller
	UserId int //用户id 当然实际上不一定是int，具体看业务要求，一般不做分表分库单表的id肯定够

}

//重写beego中的Prepare钩子函数
func (self *BaseController) Prepare() {

	//若果当前实例化的Controller是BaseControllerInterface
	//注意这里的self.AppController调用方法，必须app.CheckPermission()这样调用子类的方法，如果self.CheckPermission()
	//是没法找到子类CheckPermission的，这是因为beego的声明周期决定的
	//这样封装后只要子类可以会根据自己的业务逻辑重载对应的业务方法，注意CheckPermission返回的是bool
	//CheckValidPermission返回的是map[string]interface{}

	//千万要注意这里真的只能写成self.AppController这样，如果写成self意味着在当前下面必须实现Authorization等函数，那么在接口层面的http中封装的函数讲不会
	//被调用 然而之所我们能用钩子函数Prepare是beego特殊处理的，做不到像python那样的mro模型调用，我真他妈快调死了

	//且这行代码表示如果接口中的Controller是满足协议的化BaseControllerInterface，才会执行Authorization和CheckPermission，CheckValidPermission函数，因为认证和权限一般成对出现
	//CheckValidPermission理论上应该是不受影响的但是我没找到合适的方式封装，所以现在的使用逻辑是：要么都不用要么三个必须继承
	//当然也可以考虑把CheckValidPermission函数放到接口的控制函数中。。。。可能只是get接口的时候少一点代码把
	if app, ok := self.AppController.(BaseControllerInterface); ok {
		//认证,默认返回userid为0 不认证
		UserId, err := app.Authorization(self.Ctx)
		//如果认证解析出错 直接401
		if err != nil {
			self.Abort("401")
		} else { //否则解析当前user_id
			self.UserId = UserId
		}

		fmt.Println("userid", UserId)
		fmt.Println("CheckPermission")
		//权限校验 不通过直接403 路由权限校验
		routerStatus := app.CheckPermission(self.Ctx)
		if routerStatus == false {
			self.Abort("403")
		}
		//业务逻辑校验
		validate := app.CheckValidPermission(self.Ctx)
		if validate["code"] != 200 {
			self.Data["json"] = validate
			self.ServeJSON()
		}
	}

	return
}

//业务Controller必须全部封装下面的三个函数才算满足interface

////认证钩子函数 默认不做认证
//func (self *BaseController) Authorization(ctx *context.Context) int {
//	return 0
//}
//
////权限钩子函数 默认通过
//func (self *BaseController) CheckPermission(ctx *context.Context) bool {
//	return true
//}
//
////业务逻辑层面权限校验钩子函数，默认校验通过。比如验证码过期 您还没有权限操作
//func (self *BaseController) CheckValidPermission(ctx *context.Context) {
//
//}

// todo 还差封装不同方法的CheckValidPermission 比如get方法的时候需要校验时间格式对不对 不过看看能不能在validation那里做
// todo 因为我感觉那里更像是drf的serializer中的序列化器 具体要学到后面才能决定
