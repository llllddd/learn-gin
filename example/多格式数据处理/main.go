package main

import (
	"net/http"
)

import (
	"github.com/gin-gonic/gin"
)

type Login struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"pwd" json:"pwd" binding:"required"`
}

func main() {
	router := gin.Default()
	// Example for binding JSON ({"user":"manu","password":"123"})
	router.POST("/loginJSON", func(c *gin.Context) {
		var json Login
		if c.BindJSON(&json) == nil {
			if json.User == "lidian" && json.Password == "123" {
				c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			}
		}

	})

	router.POST("/loginForm", func(c *gin.Context) {
		var form Login
		if c.Bind(&form) == nil {
			if form.User == "lidian" && form.Password == "123" {
				c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			}
		}

	})

	router.Run(":3333")
}
