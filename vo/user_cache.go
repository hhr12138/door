package vo

import "github.com/hhr12138/door/entity"

type UserCache struct {
	User *entity.User
	CacheTimeOut int64 //过期ms值
}
