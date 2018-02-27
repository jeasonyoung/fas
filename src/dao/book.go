package dao

import (
	"errors"

	"go.uber.org/zap"

	"github.com/satori/go.uuid"

	"fas/src/log"
)

//通讯录
type Book struct {
	Id string//通讯录ID
	UserId string//用户ID

	FriendName string//朋友姓名
	FriendMobile string//朋友手机
}

//通讯录与用户关联
type UserBook struct {
	Id string//关联ID

	UserId string//用户ID
	BookId string//通讯录ID
}

//根据用户ID加载通讯录数据
func (b *Book) LoadBooksByUser(userId string)(*[]Book) {
	log.Logger.Debug("loadBooksByUser", zap.String("userId", userId))
	if len(userId) == 0 {
		log.Logger.Fatal("user Id is empty")
		return nil
	}
	if SqlDb != nil {
		log.Logger.Fatal("sql db is null")
		return nil
	}
	//
	rows, err := SqlDb.Query("select id,userId,friendName,friendMobile from tbl_fas_books where userId=? order by friendName", userId)
	if err != nil {
		log.Logger.Fatal(err.Error())
	}
	defer rows.Close()
	//
	books := make([]Book, 0)
	for rows.Next() {
		var book Book
		rows.Scan(&book.Id, &book.UserId, &book.FriendName, &book.FriendMobile)
		books = append(books, book)
	}
	if err = rows.Err(); err != nil {
		log.Logger.Fatal(err.Error())
	}
	return &books
}

//新增或更新数据
func (b *Book) SaveOrUpdate()(bool,error){
	log.Logger.Debug("saveOrUpdate")
	if SqlDb != nil {
		log.Logger.Fatal("sql db is null")
		return false, errors.New("sql db is null")
	}
	if len(b.UserId) == 0 {
		log.Logger.Fatal("userId is empty")
		return false, errors.New("userId is empty")
	}
	if len(b.Id) == 0 {//新增
		b.Id = uuid.NewV4().String()
		//
		_, err := SqlDb.Exec("insert tbl_fas_books(id,userId,friendName,friendMobile) values(?,?,?,?)",
			b.Id, b.UserId, b.FriendName, b.FriendMobile)
		if err != nil {
			log.Logger.Fatal(err.Error())
			return false, err
		}
		return true, nil
	}
	//更新
	_, err := SqlDb.Exec("update tbl_fas_books set userId=?,friendName=?,friendMobile=? where id=?",
		b.UserId, b.FriendName, b.FriendMobile, b.Id)
	if err != nil {
		log.Logger.Fatal(err.Error())
		return false, err
	}
	return true, nil
}

//根据用户ID删除数据
func (b *Book) RemoveByUser(userId string)(bool, error){
	log.Logger.Debug("RemoveByUser", zap.String("userId", userId))
	if SqlDb != nil {
		log.Logger.Fatal("sql db is null")
		return false, errors.New("sql db is null")
	}
	if len(b.UserId) == 0 {
		log.Logger.Fatal("userId is empty")
		return false, errors.New("userId is empty")
	}
	//删除数据
	_, err := SqlDb.Exec("delete from tbl_fas_books where userId=?", userId)
	if err != nil {
		log.Logger.Fatal(err.Error())
		return false, err
	}
	return true, nil
}


//新增或更新数据
func (ub *UserBook) SaveOrUpdate() (bool, error){
	log.Logger.Debug("saveOrUpdate")
	if SqlDb != nil {
		log.Logger.Fatal("sql db is null")
		return false, errors.New("sql db is null")
	}
	if len(ub.UserId) == 0 || len(ub.BookId) == 0 {
		log.Logger.Fatal("userId or bookId is empty")
		return false, errors.New("userId or bookId is empty")
	}
	//检查是否存在
	var result bool
	err := SqlDb.QueryRow("select count(0) > 0 from tbl_fas_user_books where userId=? and bookId=?", ub.UserId, ub.BookId).Scan(&result)
	if err != nil {
		log.Logger.Fatal(err.Error())
		return false, err
	}
	if result {
		return false, errors.New("userId or bookId is empty")
	}
	if len(ub.Id) == 0 {
		ub.Id = uuid.NewV4().String()
	}
	//新增
	_, err = SqlDb.Exec("insert into tbl_fas_user_books(id,userId,bookId) values(?,?,?)", ub.Id, ub.UserId, ub.BookId)
	if err != nil {
		log.Logger.Fatal(err.Error())
		return false, err
	}
	return true, nil
}

