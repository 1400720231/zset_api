package common

import "strings"
import "github.com/beego/beego/v2/server/web/context"
import "fmt"

//提供的认证函数
// jwt认证
func JwtAuthorization(ctx *context.Context) (int, error) {

	//从请求头中解析出token的参数
	authorization := ctx.Input.Header("Authorization")
	//去空格
	token := strings.TrimSpace(authorization)
	//jwt 解析token返回解析体
	user, err := ParseJwtToken(token)
	fmt.Println(user)
	//如果err有错则说明解析有问题 需要重新登陆
	if err == nil {
		fmt.Println(err)
		//return 0 //token校验失败
		ctx.Input.SetData("user_id", user.UserId)

		//self.StopRun()
	} else { //解析没有错的话把当前request.user上

		return 0, err

	}
	return user.UserId, nil
}
