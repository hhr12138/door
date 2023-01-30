package mapper

import (
	"github.com/hhr12138/door/entity"
	"github.com/hhr12138/door/object"
	"time"
)

func InsertUser(user *entity.User) (int64, error) {
	sql := "insert into `user`(`name`,`password`,`salt`,`del`,`gmt_create`,`gmt_modified`) values(?,?,?,?,?,?)"
	now := time.Now().UnixMilli()
	result, err := object.Db.Exec(sql, user.Name, user.Password, user.Salt, false, now, now)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	return id, err
}

func ExistUser(username int64, passwd string) (bool, error) {
	sql := "select count(1) from `user` where `id`=? and `password`=?"
	var result int
	err := object.Db.Get(result,sql,username,passwd)
	return result == 1,err
}

func GetSaltById(username int64) (string,error){
	sql := "select salt from `user` where `id`=?"
	var result string
	err := object.Db.Get(result,sql,username)
	return result,err
}

func GetUserById(username int64) (*entity.User, error){
	sql := "select * from `user` where `id`=?"
	result := &entity.User{}
	err := object.Db.Get(result,sql,username)
	return result,err
}
