package routes

import (
	"douyin-backend/go/controller"
	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {
	r := gin.Default()

	/**
	  用户User路由组
	*/
	userGroup := r.Group("/douyin/user")
	{
		//注册
		userGroup.POST("/register", controller.Register)
		//登录
		userGroup.POST("/login", controller.Login)
		//用户信息
		userGroup.POST("/", controller.UserInfo)

		////修改某个User
		//userGroup.PUT("/users/:id", controller.UpdateUser)
		////删除某个User
		//userGroup.DELETE("/users/:id", controller.DeleteUserById)
	}

	return r
}
