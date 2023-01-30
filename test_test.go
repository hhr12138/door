package door

import (
	"github.com/gin-gonic/gin"
	"github.com/hhr12138/door/object"
	"github.com/jmoiron/sqlx"
	"testing"
)

func TestRegister(t *testing.T) {
	db,_ := sqlx.Open("mysql",object.TEST_DB_ADDR)
	engine := gin.Default()
	Register(db,engine)
	engine.Run()
}
