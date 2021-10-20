package topic

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web/context"
	"zset_api/common"
	"zset_api/models/topic"
)

type AllTopicController struct {
	common.BaseController
}

func (self *AllTopicController) Get() {
	data := []map[string]interface{}{}                                         //声明一个不定长的map类型的数组
	response := map[string]interface{}{"code": 200, "message": "", "data": ""} //默认状态码200

	o := orm.NewOrm()
	//注意这里的写法只声明变量类型 .All()源码中是这么写的。
	var topics []*topic.Topic

	//select id name from topic;
	count, _ := o.QueryTable("topic").Filter("is_delete", false).Filter("is_active", true).All(&topics, "Id", "Name")

	fmt.Println(count)
	response["code"] = 200
	response["message"] = "success"
	//迭代序列化数据封装
	for _, value := range topics {
		data = append(data, value.GetTopic())
	}
	response["data"] = data
	self.Data["json"] = response
	self.ServeJSON()
}
func (self *AllTopicController) Authorization(ctx *context.Context) (int, error) {
	return 0, nil
}

// 认通过
func (self *AllTopicController) CheckPermission(ctx *context.Context) bool {
	return true
}

// CheckValidPermission form表单校验
func (self *AllTopicController) CheckValidPermission(ctx *context.Context) map[string]interface{} {
	response := map[string]interface{}{"code": 200, "message": "", "data": ""} //默认状态码200

	return response
}
