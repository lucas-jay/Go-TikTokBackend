package main

import (
	"github.com/gin-gonic/gin"
	"testProject/common"
)

func main() {

	//获取初始化的数据库
	db := common.InitDB()

	defer db.Close()

	//创建一个默认的路由引擎
	r := gin.Default()

	//启动路由
	CollectRoutes(r)

	//在9090端口启动服务
	panic(r.Run(":9090"))
}
