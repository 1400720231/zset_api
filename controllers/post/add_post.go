package post

import (
	"github.com/beego/beego/v2/server/web/context"
	"zset_api/common"
)

type AddPostController struct {
	common.BaseController
}

//form
type addPostForm struct {
}

//json application/json
type addPostJson struct {
}

func (self *AddPostController) Post() {
	var data []map[string]interface{}                                             //type reference 声明一个不定长的map类型的数组
	var response = map[string]interface{}{"code": 200, "message": "", "data": ""} //默认状态码200

	response["data"] = data
	self.Data["json"] = response
	self.ServeJSON()
}
func (self *AddPostController) Authorization(ctx *context.Context) (int, error) {
	userId, err := common.JwtAuthorization(ctx)

	return userId, err
}

// 认通过
func (self *AddPostController) CheckPermission(ctx *context.Context) bool {
	ok := common.IsLoginPermission(ctx)
	if ok {
		return true
	}
	return false
}

// CheckValidPermission form表单校验
func (self *AddPostController) CheckValidPermission(ctx *context.Context) map[string]interface{} {
	response := map[string]interface{}{"code": 200, "message": "", "data": ""} //默认状态码200

	return response
}
