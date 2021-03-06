package user

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

//如果不设置colume字段 数据库字段就是小写的username
//model的定义见文档  https://beego.me/docs/mvc/model/models.md
type User struct {
	Id         int       `orm:"pk;auto"`                                             //id 自增
	Username   string    `orm:"unique;column(user_name);size(64);description(用户名);"` //用户名 数据库字段会是user_name
	Password   string    `orm:"size(32);description(密码)"`                            //json方式序列化的时候舍弃这个字段不返回
	Age        int       `orm:"null;description(年龄)"`
	Gender     int       `orm:"null;description(1:男,2:女,3:未知)"`
	Phone      int64     `orm:"null;description(电话号码)"`
	Email      string    `orm:"null;size(255);description(邮箱)"`
	IsActive   int       `orm:"description(1启用，0停用);default(1)"`  //默认值是1
	IsDelete   int       `orm:"description(1删除，0未删除);default(0)"` //默认值是0
	Profile    string    `orm:"null;size(255);description(头像oss)"`
	CreateTime time.Time `orm:"auto_now_add;type(datetime);description(创建时间);null"` //创建时间自动添加 datetime类型；也有date类型选择
	UpdateTime time.Time `orm:"auto_now;type(datetime);description(更新时间);null"`     //修改时间自动添加 每次修改都会自己变 datetime类型也有date类型选择
}

//指定表在数据库中的名字
func (u *User) TableName() string {
	return "user"
}

//序列化方法 注意这里一定要大写！！！不然不能访问 我日你妈！！！
//可以为不同的接口定义不同的序列化数据 还可以封装函数满足类似drf serializer method的功能 绝了
func (u User) Userinfo() map[string]interface{} {
	user := make(map[string]interface{})
	user["id"] = u.Id
	user["username"] = u.Username
	user["age"] = u.Age
	user["gender"] = u.Gender
	user["phone"] = u.Phone
	user["is_active"] = u.IsActive
	user["profile"] = u.Profile
	user["create_time"] = u.CreateTime
	return user
}

//orm注册表结构 不然调用syncdb的时候不会识别这个model 也就不会迁移成功
func init() {
	orm.RegisterModel(
		new(User),
	)
}
