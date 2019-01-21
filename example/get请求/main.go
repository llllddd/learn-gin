package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/user", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "li")
		lastname := c.Query("lastname")
		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	router.Run(":3333")
}
