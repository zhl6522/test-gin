package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"test-gin/test-file/model"
)

// 绑定JSON的示例 ({"user": "zhl", "password": "123456"})
// 绑定form表单示例 (user=q1mi&password=123456)
// 绑定QueryString示例 (/loginQuery?user=q1mi&password=123456)
func LoginUser(c *gin.Context) {
	var login model.Login
	if err := c.ShouldBind(&login); err ==nil {
		fmt.Printf("login info:%#v\n", login)
		c.JSON(http.StatusOK, gin.H{
			"user":login.User,
			"password":login.Password,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"err":err.Error()})
	}
}

func File(c *gin.Context) {
	c.HTML(http.StatusOK, "upload/index.html", gin.H{})
}

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("f1")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":err.Error(),
		})
		return
	}
	log.Println(file.Filename)
	dst := fmt.Sprintf("E:/www/go_project/src/test-gin/uploads/%s", file.Filename)
	// 上传文件到指定的目录
	c.SaveUploadedFile(file, dst)
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("'%s' uploaded!", file.Filename),
	})
}

func Files(c *gin.Context) {
	c.HTML(http.StatusOK, "upload/indexs.html", gin.H{})
}

func UploadFiles(c *gin.Context) {
	form,_ := c.MultipartForm()
	files := form.File["upload[]"]
	for _, file := range files {
		log.Println(file.Filename)
		dst := fmt.Sprintf("E:/www/go_project/src/test-gin/uploads/%s", file.Filename)
		//上传文件到指定的目录
		c.SaveUploadedFile(file,dst)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%d files uploaded!", len(files)),
	})
}