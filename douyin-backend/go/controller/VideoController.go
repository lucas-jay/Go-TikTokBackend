package controller

import (
	"fmt"
	"github.com/disintegration/imaging"
	//gin框架的依赖
	"github.com/gin-gonic/gin"
	//http连接包
	"net/http"
)

type PublishInfoResponse struct {
	status_code int32
	status_msg  string
}

func PostVideo(c *gin.Context) {

	file, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//上传file文件到指定的文件路径
	path := "111"
	video_name := "111" + ".mp4"
	fn := path + video_name
	if err := c.SaveUploadedFile(file, fn); err != nil {
		fmt.Println("存储视频失败，filePath,", fn)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//读取视频第一帧作为封面
	//reader := utils.ReadFrameAsJpeg(fn)
	img, err := imaging.Decode(reader)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	image_path := "222"

	// replace .mp4 to .jpg
	image_name := video_name[:len(video_name)-4] + ".jpg"
	url := image_path + image_name
	if err = imaging.Save(img, url); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//把它们给存到数据库里面
	//play_url := video_name
	//cover_url := image_name

	c.JSON(http.StatusOK, PublishInfoResponse{
		status_code: 0,
		status_msg:  "ok",
	})

	//todo
	//传过来的token没接，要鉴权，已登陆？
	//视频数据库没建
	//视频封面处理还没做
	//视频保存路径，封面保存路径，video_name
	//白天再继续完成这些吧

}

/**
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
*/
