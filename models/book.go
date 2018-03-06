package models

import "github.com/astaxie/beego/orm"

//通讯录
type Book struct {
	Id string `orm:"column(id);pk"`//通讯录ID
	UserId string `orm:"column(userId)"`//所属用户ID

	FriendName string `orm:"column(friendName)"`//朋友姓名
	FriendMobile string `orm:"column(friendMobile)"`//朋友手机
}

//通讯录-表名
func (b *Book) TableName() string {
	return "tbl_fas_books"
}

//通讯录与用户关联
type UserBook struct {
	Id string `orm:"column(id);pk"`//关联ID

	UserId string `orm:"column(userId)"`//用户ID
	BookId string `orm:"column(bookId)"`//通讯录ID
}

//通讯录与用户关联-表名
func (ub *UserBook) TableName() string {
	return "tbl_fas_user_books"
}


//注册表
func init(){
	orm.RegisterModel(new(Book), new(UserBook))
}