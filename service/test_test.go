package service

import (
	"fmt"
	"github.com/hhr12138/door/object"
	"log"
	"testing"
)


func init() {
	object.RegisterTest()
}

func TestExitUser(t *testing.T) {
	username := int64(10000000020)
	passwd := "hhr12138"
	request := &LoginRequest{
		Username: username,
		Password: passwd,
	}
	reply, err := ExistUser(request)
	if err != nil{
		log.Fatalln(err)
	}
	if reply.Success{
		fmt.Println("token=:"+reply.Token)
	}
}

func TestRegister(t *testing.T) {
	request := &RegisterRequest{
		Name:"我爱上厕所",
		Password: "hhr12138",
		Salt: "qwer",
	}
	username, err := Register(request)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(username)
}