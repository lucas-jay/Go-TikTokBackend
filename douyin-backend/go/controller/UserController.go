package controller

import (
	//需要用到的结构体
	"douyin-backend/go/entity"
	//gin框架的依赖
	"github.com/gin-gonic/gin"
	//http连接包
	"net/http"
	//service层方法
	"douyin-backend/go/service"
)

func Register(c *gin.Context) {
	//定义一个User变量
	var user entity.User
	//将调用后端的request请求中的body数据根据json格式解析到User结构变量中
	err := c.BindJSON(&user)
	if err != nil {
		return
	}
	//将被转换的user变量传给service层的CreateUser方法，进行User的新建
	err = service.CreateUser(&user)
	//判断是否异常，无异常则返回包含200和更新数据的信息
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": user,
		})
	}
}

func Login(c *gin.Context) {

	var user entity.User
	user.Name = c.Query("username")
	user.Password = c.Query("password")

	user, _ = service.GetUserByName(user.Name)

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"msg":     2,
		"user_id": user.Id,
		"token":   4,
	})

}

func UserInfo(c *gin.Context) {
	var user entity.User
	user, _ = service.GetUserByName("Lucas")

	//user_id := c.Query("user_id")
	//token := c.Query("token")

	c.JSON(http.StatusOK, UserInfoResponse{
		status_code: 0,
		status_msg:  "ok",
		user:        user,
	})
}

type UserInfoResponse struct {
	status_code int32
	status_msg  string
	user        entity.User
}

func GetUserList(c *gin.Context) {
	todoList, err := service.GetAllUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": todoList,
		})
	}
}
