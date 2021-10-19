package user

import (
	"github.com/beego/beego/v2/server/web/context"
	"zset_api/common"
)

type UploadController struct {
	common.BaseController
}

//https://beego.me/docs/mvc/controller/params.md 文件处理
//利用结构体的绑定实现http的post方法 beego内部对应的Post函数
func (self *UploadController) Post() {
	data := make(map[string]interface{})
	response := map[string]interface{}{"code": 200, "message": "文件上传成功", "data": ""} //默认状态码200
	f, _, err := self.GetFile("file")
	if err != nil {
		response["code"] = 600
		response["message"] = "文件上传出错"
	}
	link := common.UploadFileToOss(f, "go封面.jpeg")

	data["message"] = link
	response["data"] = data
	self.Data["json"] = response
	defer f.Close()
	self.ServeJSON()
}

//认证
func (self *UploadController) Authorization(ctx *context.Context) (int, error) {

	return 0, nil

}

// 权限
func (self *UploadController) CheckPermission(ctx *context.Context) bool {

	return true
}

//业务校验
func (self *UploadController) CheckValidPermission(ctx *context.Context) map[string]interface{} {
	return map[string]interface{}{"code": 200}
}
