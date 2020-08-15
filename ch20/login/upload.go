package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	//"log"
	"net/http"
)

func main() {
	r := gin.Default()
	// 限制表单上传大小 2M
	r.MaxMultipartMemory = 2 << 20
	r.POST("/upload", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get err %s", err.Error()))
		 	return
		}
		// 获取所有图片
		files := form.File["files"]
		// 遍历所有图片
		for _,file := range files {
			// 依次存
			if err :=c.SaveUploadedFile(file, file.Filename);err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("upload err %s", err.Error()))
				return
			}
		}
		c.String(http.StatusOK, fmt.Sprintf("upload ok %d file!", len(files)))

		/*// 表单提交
		file,_ := c.FormFile("file")
		log.Println(file.Filename)
		// 传到项目根目录，名字就用本身的
		c.SaveUploadedFile(file, file.Filename)
		// 打印信息
		c.String(http.StatusOK, fmt.Sprintf("'%s' upload!", file.Filename))*/
	})
	r.Run()
}
