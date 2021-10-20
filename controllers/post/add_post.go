package post

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/validation"
	"github.com/beego/beego/v2/server/web/context"
	"strings"
	"zset_api/common"
	"zset_api/models/post"
	"zset_api/models/topic"
	"zset_api/models/user"
)

type AddPostController struct {
	common.BaseController
}

//form
//validation: https://beego.me/docs/mvc/controller/validation.md#%E7%A4%BA%E4%BE%8B
//Required不传参校验不通过ok=false，不会抛出错误err=nil
//Min(1)表示最小值是1 ，如果值是0，ok=false不会抛出错误err=nil，
//如果topic="大大说"不是一个int才会抛出错误err
//也就是说判断一个表单是否校验通过其核心是看ok是否为true，err是不满足类型格式校验的时候才会有值
type addPostForm struct {
	TopicId  int    `form:"topic_id" valid:"Required;Min(1)"` //
	Title    string `form:"title" valid:"Required;MinSize(1);MaxSize(50)"`
	Content  string `form:"content" valid:"Required;MinSize(1);MaxSize(5000)"`
	IsActive int    `form:"is_active" valid:"Required;Min(1)"`
}

//json application/json
type addPostJson struct {
	TopicId  string `form:"topic_id"`
	Title    string `form:"title"`
	Content  string `form:"content"`
	IsActive int    `form:"is_active"`
}

func (self *AddPostController) Post() {
	var response = map[string]interface{}{"code": 200, "message": "", "data": ""} //默认状态码200
	err, message := self.Validate(self.Ctx)

	if err != nil {
		response["code"] = 600
		response["message"] = message
		self.Data["json"] = response
		self.ServeJSON()
		//注意这里必须要return或者self.StopRun()，不然函数会继续往if结构体后面执行
		//通过源码可知道，ServeJSON只是设置了上下文中http response中的headers：application/json,然后跟你设置的self.Data["json"] = response
		//判断一下是否能用json格式返回，并返回一个error类型。
		//所以ServeJSON并不是Post方法的结束标志，还会往下执行，所以你需要手动return或者self.StopRun()

		//但是我同时发现了另外一个问题似乎ServeJSON在整个证明周期中只对第一次执行的内容生效
		//后面再次执行ServeJSON不会修改返回的内容self.Data["json"]，不知道为什么

		//return
		//self.StopRun()
	}

	fmt.Println("ServeJSON")
	o := orm.NewOrm()
	topic_id, _ := self.GetInt("topic_id")
	is_active, _ := self.GetInt("is_active")
	//方式1
	topic_obj := &topic.Topic{Id: topic_id}
	user_obj := &user.User{Id: self.UserId}
	//方式2
	//topic_obj := &topic.Topic{}
	//user_obj := &user.User{}
	//o.QueryTable("topic").Filter("Id", topic_id).One(topic_obj)
	//o.QueryTable("user").Filter("Id", self.UserId).One(user_obj)
	post_obj := post.Post{
		UserId:   user_obj, //注意这里是个user的结构体实例化对象，而post只需要这个user结构体中的id，所以你直接实例化或者orm查询都行
		TopicId:  topic_obj,
		Content:  self.GetString("content"),
		Title:    self.GetString("title"),
		IsActive: is_active,
	}
	o.Insert(&post_obj)
	response["data"] = map[string]interface{}{"id": post_obj.Id}
	self.Data["json"] = response
	response["message"] = "发帖成功"
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

//  form表单校验基本校验 相当于接口文档层面校验
func (self *AddPostController) CheckValidPermission(ctx *context.Context) map[string]interface{} {
	response := map[string]interface{}{"code": 200, "message": "", "data": ""} //默认状态码200
	form := addPostForm{}
	err := self.ParseForm(&form)

	valid := validation.Validation{}
	//u := addPostForm{TopicId: 0}
	if err == nil {
		//Required不传参校验不通过ok=false，不会抛出错误err=nil
		//Min(1)表示最小值是1 ，如果值是0，ok=false不会抛出错误err=nil，
		//如果topic="大大说"不是一个int才会抛出错误err
		//也就是说判断一个表单是否校验通过其核心是看ok是否为true，err是不满足类型格式校验的时候才会有值
		ok, _ := valid.Valid(&form)
		if ok != true {
			response["code"] = 400
			response["message"] = "参数错误"
			response["data"] = form
		}
	}
	return response
}

// 表单的业务逻校验
func (self *AddPostController) Validate(ctx *context.Context) (error, string) {
	content := self.GetString("content", "")
	index := strings.Index(content, "sb")
	if index != -1 {
		return fmt.Errorf("校验失败"), "包含敏感词汇"
	} else {
		return nil, ""
	}
}
