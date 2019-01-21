package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		test := c.DefaultPostForm("name", "llli")

		c.JSON(http.StatusOK, gin.H{"status": gin.H{"status_code": http.StatusOK, "status": "ok"}, "message": message, "name": test})
	})
	router.Run(":3333")
}

//curl -X POST http://localhost:3333/form_post -H "Content-Type:application/x-www-form-urlencoded" -d "message=hello&name=aaa"
