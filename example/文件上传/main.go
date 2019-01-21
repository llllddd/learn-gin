package main

import (
	"fmt"
)

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/upload", func(c *gin.Context) {
		//获取页面name标签的名字
		name := c.PostForm("name")
		fmt.Println(name)
		//获取文件
		file, err := c.FormFile("upload")
		if err != nil {
			c.String(http.StatusBadRequest, "a Bad request")
			return
		}
		//获取文件名
		filename := file.Filename
		fmt.Println("=================", filename)
		if err := c.SaveUploadedFile(file, file.Filename); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err %s", err.Error()))
			return
		}
		c.String(http.StatusBadRequest, "upload successful")
		c.String(http.StatusCreated, "upload successful")
	})

	router.Run(":3333")
}
