package service

import (
	"database/sql"
	"github.com/hhr12138/chat_room-utils/encrypt"
	"github.com/hhr12138/door/entity"
	"github.com/hhr12138/door/mapper"
	"github.com/hhr12138/door/object"
)

type RegisterRequest struct {
	Name     string `from:"name",binding:"required"`
	Password string `from:"password",binding:"required"`
	Salt     string `from:"salt"`
}

type RegisterReply struct {
	Username int64 `json:"username"`
}

type LoginRequest struct {
	Username int64  `from:"username",binding:"required"`
	Password string `from:"password",binding:"required"`
}

type LoginReponse struct {
	Success bool
	Token   string
}

type LogoutRequest struct {
	Token string `from:"token",binding:"required"`
}

func Register(request *RegisterRequest) (*RegisterReply, error) {
	passwd := encrypt.EncryptPassword(request.Password, request.Salt)
	user := &entity.User{
		Name: sql.NullString{
			String: request.Name,
			Valid:  true,
		},
		Password: sql.NullString{
			String: passwd,
			Valid:  true,
		},
		Salt: sql.NullString{
			String: request.Salt,
			Valid:  true,
		},
	}
	username, err := mapper.InsertUser(user)
	reply := &RegisterReply{
		Username: username,
	}
	return reply, err
}

func ExistUser(request *LoginRequest) (*LoginReponse, error) {
	username := request.Username
	reply := &LoginReponse{}
	user, err := mapper.GetUserById(username)
	if err != nil {
		return nil, err
	}
	passwd := encrypt.EncryptPassword(request.Password, user.Salt.String)
	if !user.Password.Valid || passwd != user.Password.String{
		return reply,nil
	}
	reply.Success = true
	token := encrypt.Token(username, passwd)
	reply.Token = token
	object.CacheUser(token,user)
	return reply, nil
}

