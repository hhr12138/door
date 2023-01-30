package object

import (
	"database/sql"
	"fmt"
	"github.com/hhr12138/door/entity"
	"strconv"
	"testing"
	"time"
)

var token string

func init() {
	token = strconv.FormatInt(time.Now().UnixMilli(),10)
	RegisterTest()
}

func TestCacheUser(t *testing.T) {
	CacheUser(token, &entity.User{
		Name: sql.NullString{
			String: "lisi",
			Valid:  true,
		},
		Password: sql.NullString{
			String: "123456789",
			Valid:  true,
		},
	})
}

func TestGetUser(t *testing.T) {
	user, ok := GetUser("55c9c9552904f45935de18bc298d3978")
	if ok {
		fmt.Println(user.Name.String+" "+user.Password.String)
	}
}

func TestRemoveUser(t *testing.T) {
	RemoveUser(token)
}
