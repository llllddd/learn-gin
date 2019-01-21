package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"time"
)

type Person struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func main() {
	router := gin.Default()
	router.GET("/testing", startPage)
	router.Run(":3333")

}

func startPage(c *gin.Context) {
	var person Person

	if c.ShouldBind(&person) == nil {
		log.Println(person.Name)
		log.Println(person.Address)
		log.Println(person.Birthday)
	}

	c.String(200, "Success")
}
