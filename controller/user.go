package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hhr12138/door/service"
	"net/http"
)

func Register(ctx *gin.Context) {
	request := new(service.RegisterRequest)
	err := ctx.ShouldBind(request)
	if err != nil{
		ctx.JSON(http.StatusBadGateway,err)
		return
	}
	reply, err := service.Register(request)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError,err)
	} else {
		ctx.JSON(http.StatusOK,reply)
	}
}
