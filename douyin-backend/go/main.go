package main

import (
	"douyin-backend/go/dao"
	"douyin-backend/go/entity"
	"douyin-backend/go/routes"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	//连接数据库
	err := dao.InitMySql()
	if err != nil {
		panic(err)
	}
	//程序退出关闭数据库连接
	defer dao.Close()
	//绑定模型
	dao.SqlSession.AutoMigrate(&entity.User{})
	//注册路由
	r := routes.SetRouter()
	//启动端口为8085的项目
	r.Run(":8081")
}
