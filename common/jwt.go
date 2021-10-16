package common

import (
	beego "github.com/beego/beego/v2/server/web"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

//从配置文件中获取SECRET_KEY 类似所有后台系统加密的key
var secret_key, _ = beego.AppConfig.String("SECRET_KEY")

//
var jwtSecret = []byte(secret_key)

//你打算用来放在加密内容中的东西，比如django放了username
type JwtPayload struct {
	Username string `json:"username"` //用户名（假设用户名唯一）
	UserId   int    `json:"userid"`   //用户唯一标示
	jwt.StandardClaims
}

// 生成jwt token
func GenerateJwtToken(username string, userid int) (string, error) {
	//当前时间
	nowTime := time.Now()
	//token的到期时间 默认7天
	expireTime := nowTime.Add(60 * 24 * 7 * time.Second)

	claims := JwtPayload{
		username,
		userid,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //把datetime时间转化成时间戳
		},
	}
	//hash256方式加密
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//加盐内容是我们的配置的secret_key 和django或者说和大多数的加密方式一样
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// 解析验证jwt token
func ParseJwtToken(token string) (*JwtPayload, error) {
	//匿名函数 没有函数名 声明并立即调用
	tokenClaims, err := jwt.ParseWithClaims(token, &JwtPayload{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	//如果解析成功
	if tokenClaims != nil {
		//有jwt 解析成功 且没有过期
		if claims, ok := tokenClaims.Claims.(*JwtPayload); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	//如果解析失败(没有传jwt值参数) 或者 解析失败（解析错误和过期）
	return nil, err
}

/*
//传入用户名和系统中的userid标示
	token, _ := common.GenerateToken("xiongyao", "431")
	fmt.Println(token)
	ParseToken, err := common.ParseToken(token)
	if err!=nil{
		fmt.Println("token解析失败 请重新登陆")
	}
	//获取当前的username userid
	fmt.Println(ParseToken.Username,ParseToken.UserId)
*/
