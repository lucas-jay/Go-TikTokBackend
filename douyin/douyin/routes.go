package main

import (
	"testProject/controller"

	"github.com/gin-gonic/gin"
)

func CollectRoutes(r *gin.Engine) *gin.Engine {

	//注册业务Api
	r.POST("/douyin/user/register/", controller.Register)
	//登录业务Api
	r.POST("/douyin/user/login/", controller.Login)

	return r

}
