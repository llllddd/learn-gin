package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//:用来指定路由参数，*也可以用来代替相关参数
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is" + action
		c.String(http.StatusOK, message)
	})
	router.Run(":3333")
}
