package main

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	_ "zset_api/routers"
	//一定要倒入这个包 不然各种奇怪的错误 比如无法迁移model
	_ "github.com/go-sql-driver/mysql"
	//把对应的model的package中的init函数给执行了！！！
	_ "zset_api/models/post"
	_ "zset_api/models/topic"
	_ "zset_api/models/user"
)

func init() {
	logs.SetLogger(logs.AdapterMultiFile, `{"filename":"logs/log.out","separate":["error","info"]}`)
	//beego2返回两个参数，对参数名找不到的情况做了错误处理
	username, _ := beego.AppConfig.String("username")
	pwd, _ := beego.AppConfig.String("pwd")
	host, _ := beego.AppConfig.String("host")
	port, _ := beego.AppConfig.String("port")
	db, _ := beego.AppConfig.String("db")

	// root:250onioN!!!!@tcp(localhost:3306)/zset?charset=utf8
	dataSource := username + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + db + "?charset=utf8"
	//mysql引擎
	err1 := orm.RegisterDriver("mysql", orm.DRMySQL)
	//这里的"default"相当于一个别名，万一一个是主一个是从数据库，需要读写分离的时候就可以根据这个做映射
	err2 := orm.RegisterDataBase("default", "mysql", dataSource)
	logs.Info(dataSource)
	logs.Info(err1, err2)
	//\n可以让输出的时候没有%出现
	fmt.Printf("host:%s|port:%s|db:%s\n", host, port, db)

}

func main() {
	//先执行init函数 再执行main函数 说明init是一个package中最先执行的函数，优先级大于main

	//数据库命令行迁移，接受go run main [orm] options
	orm.RunCommand()
	//直接执行数据库迁移操作
	//err := orm.RunSyncdb("default", false, true)
	//logs.Info(err)
	orm.Debug = true

	beego.Run()

}
