package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hhr12138/door/object"
	"github.com/hhr12138/door/service"
	"net/http"
)

func Login(ctx *gin.Context){
	request := new(service.LoginRequest)
	err := ctx.ShouldBind(request)
	if err != nil{
		ctx.JSON(http.StatusBadGateway,err)
		return
	}
	reply, err := service.ExistUser(request)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError,err)
	} else{
		ctx.JSON(http.StatusOK,reply)
	}
}

func Logout(ctx *gin.Context) {
	request := new(service.LogoutRequest)
	err := ctx.ShouldBind(request)
	if err != nil{
		ctx.JSON(http.StatusBadGateway,err)
	}
	object.RemoveUser(request.Token)
	ctx.JSON(http.StatusOK,"Success")
}

