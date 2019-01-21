package main

import (
	"log"
)

import (
	"github.com/gin-gonic/gin"
)

type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

func main() {
	router := gin.Default()
	router.Any("/testing", test)
	router.Run(":3333")
}

func test(c *gin.Context) {
	var person Person
	if c.ShouldBindQuery(&person) == nil {
		log.Println("=====Only Bind By Query String =====")
		log.Println(person.Name)
		log.Println(person.Address)
	}
	c.String(200, "Success")
}
