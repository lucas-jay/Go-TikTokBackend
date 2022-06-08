package controller

import (
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"io"
	"os"
	"strconv"
	"time"

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

	//接token
	//var user entity.User
	//user.Name = c.Query("username")
	//user.Password = c.Query("password")
	//
	//user, _ = service.GetUserByName(user.Name)

	//接title,这里接了后面没用到，是存到数据库的，数据库还没设计，所以我就先注释了
	/*
		title, ok := c.GetPostForm("title")
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	*/
	//一个不知道我们要不要的异常处理，大概是测试路径的？
	_, err = os.Executable()
	if err != nil {
		panic(err)
	}

	//上传file文件到指定的文件路径
	//path，video_name
	path := "111"
	video_name := GenerateFilename("ycc", 111) + ".mp4"
	fn := path + video_name
	if err := c.SaveUploadedFile(file, fn); err != nil {
		fmt.Println("存储视频失败，filePath,", fn)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//读取视频第一帧作为封面
	reader := ReadFrameAsJpeg(fn)
	img, err := imaging.Decode(reader)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//image_path
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
	//title

	c.JSON(http.StatusOK, PublishInfoResponse{
		status_code: 0,
		status_msg:  "ok",
	})

	//todo
	//传过来的token不知道怎么解析
	//视频数据库没建,所以数据保存还没做
	//视频保存路径，封面保存路径，video_name,
	//错误处理有点不太懂
	//以及别人的不太清楚怎么改一下，直接用是不是不太好，目前直接照搬过来了，不过我去看了一下别人用到的库
	//而且还没测试这一块
}

// GenerateFilename depend on timestamp, username and user id to generate a filename for uploaded videos
func GenerateFilename(username string, userId int64) string {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	return timestamp + "_" + username + "_" + strconv.FormatInt(userId, 10)
}

// ReadFrameAsJpeg use ffmpeg read first frame of video as a jpeg
func ReadFrameAsJpeg(inFileName string) io.Reader {
	frameNum := 1
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(inFileName).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		panic(err)
	}
	return buf
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
