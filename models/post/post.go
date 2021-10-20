package post

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
	"zset_api/models/topic"
	"zset_api/models/user"
)

//如果不设置colume字段 数据库字段就是小写的username
//model的定义见文档  https://beego.me/docs/mvc/model/models.md
type Post struct {
	Id         int          `orm:"pk;auto"`                                         //id 自增
	UserId     *user.User   `orm:"rel(fk);description(用户外键);on_delete(do_nothing)"` //外键 集联删除不操作
	TopicId    *topic.Topic `orm:"rel(fk);description(话题外键);on_delete(do_nothing)"` //外键 集联删除不操作
	Title      string       `orm:"size(64);description(帖子名字)"`
	Content    string       `orm:"type(text);description(帖子内容)"`                       //文章内容比较大 设置为text对应mysql longtext
	IsActive   int          `orm:"description(1草稿箱，0非草稿箱);default(1)"`                 //默认值是1
	IsDelete   int          `orm:"description(1删除，0未删除);default(0)"`                   //默认值是0
	CreateTime time.Time    `orm:"auto_now_add;type(datetime);description(创建时间);null"` //创建时间自动添加 datetime类型；也有date类型选择
	UpdateTime time.Time    `orm:"auto_now;type(datetime);description(更新时间);null"`     //修改时间自动添加 每次修改都会自己变 datetime类型也有date类型选择

}

//指定表在数据库中的名字
func (u *Post) TableName() string {
	return "post"
}
func (u *Post) GetPost() map[string]interface{} {
	post := make(map[string]interface{})
	post["id"] = u.Id
	post["user_id"] = u.UserId
	post["title"] = u.Title
	post["content"] = u.Content
	post["create_time"] = u.CreateTime
	post["update_time"] = u.UpdateTime
	return post
}

//orm注册表结构 不然调用syncdb的时候不会识别这个model 也就不会迁移成功
func init() {
	orm.RegisterModel(
		new(Post),
	)
}
