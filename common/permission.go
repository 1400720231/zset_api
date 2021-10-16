package common

import (
	"fmt"
	"github.com/beego/beego/v2/server/web/context"
)

//提供的权限函数
// jwt认证
func IsLoginPermission(ctx *context.Context) bool {
	//c.CruSession.Set(context2.Background(), name, value)
	//从请求头中解析出token的参数

	user_id := ctx.Input.GetData("user_id")
	fmt.Println("IsLoginPermission", user_id)
	if user_id == 0 {
		return false
	}
	return true
}
