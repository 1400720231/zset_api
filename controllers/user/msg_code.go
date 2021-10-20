package user

import (
	"github.com/beego/beego/v2/server/web/context"
	"zset_api/common"
)

type MsgController struct {
	common.BaseController
}

//https://beego.me/docs/mvc/controller/params.md 文件处理
//利用结构体的绑定实现http的post方法 beego内部对应的Post函数
func (self *MsgController) Post() {
	data := make(map[string]interface{})

	self.Data["json"] = data

	self.ServeJSON()
}

//认证
func (self *MsgController) Authorization(ctx *context.Context) (int, error) {

	return 0, nil

}

// 权限
func (self *MsgController) CheckPermission(ctx *context.Context) bool {

	return true
}

//业务校验
func (self *MsgController) CheckValidPermission(ctx *context.Context) map[string]interface{} {
	return map[string]interface{}{"code": 200}
}
