package door

import (
	"github.com/gin-gonic/gin"
	"github.com/hhr12138/door/controller"
	"github.com/hhr12138/door/filter"
	"github.com/hhr12138/door/object"
	"github.com/jmoiron/sqlx"
)

func Register(db *sqlx.DB, engine *gin.Engine) {
	object.RegisterDb(db)
	object.RegisterCache()
	registerRoute(engine)
}

func registerRoute(engine *gin.Engine) {
	//get
	engine.DELETE("/user/logout",filter.LoginFilter,controller.Logout)
	//post
	engine.POST("/user/login",controller.Login)
	engine.POST("/user/register",controller.Register)
}
