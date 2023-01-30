package object

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const(
	TEST_DB_ADDR = "root:hhr12138@(localhost)/common"
)

var Db *sqlx.DB


func RegisterDb(db *sqlx.DB){
	Db = db
}

func RegisterTest(){
	RegisterCache()
	db,_ := sqlx.Open("mysql",TEST_DB_ADDR)
	RegisterDb(db)
}