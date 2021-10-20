package post

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web/context"
	"zset_api/common"
	"zset_api/models/post"
)

type AllPostController struct {
	common.BaseController
}

func (self *AllPostController) Get() {
	var data []map[string]interface{}                                          //type reference 声明一个不定长的map类型的数组
	response := map[string]interface{}{"code": 200, "message": "", "data": ""} //默认状态码200
	//注意这里的写法只声明变量类型 .All()源码中是这么写的。
	var posts []*post.Post
	o := orm.NewOrm()

	//select id name from topic;
	count, _ := o.QueryTable("post").Filter("is_delete", 0).Filter("is_active", 1).All(&posts, "Id", "Title", "Content", "CreateTime", "UpdateTime")

	fmt.Println("count", count)
	response["code"] = 200
	response["message"] = "success"
	//迭代序列化数据封装
	for _, value := range posts {
		data = append(data, value.GetPost())
	}
	response["data"] = data
	self.Data["json"] = response
	self.ServeJSON()
}
func (self *AllPostController) Authorization(ctx *context.Context) (int, error) {
	return 0, nil
}

// 认通过
func (self *AllPostController) CheckPermission(ctx *context.Context) bool {
	return true
}

// CheckValidPermission form表单校验
func (self *AllPostController) CheckValidPermission(ctx *context.Context) map[string]interface{} {
	response := map[string]interface{}{"code": 200, "message": "", "data": ""} //默认状态码200

	return response
}
