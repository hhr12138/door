package filter

import (
	"github.com/gin-gonic/gin"
	"github.com/hhr12138/door/object"
	"net/http"
)

func LoginFilter(ctx *gin.Context){
	login := false
	token := ctx.GetHeader("token")
	if len(token) != 0{
		user, _ := object.GetUser(token)
		ctx.Set("user",user)
		login = true
	}
	if !login{
		ctx.AbortWithStatusJSON(http.StatusFound,"please login")
	}
}