package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

/*******一般写法***********************
func GetDescription(c *gin.Context) {
	resp := map[string]string{"hello": "world"}
	c.JSON(200, resp)
}
func main() {
	r := gin.Default()
	r.GET("/description", GetDescription)
	r.Run(":3333")
}
***************************************/
func GetDescription(c *gin.Context) {
	fmt.Println("I'm lidian!")
	c.Next()
}

func main() {
	r := gin.New()
	r.Use(GetDescription)
	r.Run
	(":3333")

}

/****************中间件的第二种写法**********
func GetDescription()gin.HandlerFunc{
     return func(c *gin.Context){
	     fmt.Println("I'm middle too")
     }
}

func main(){
	r := gin.New()
	r.Use(GetDescription())
	r.Run(":3333")
}
*/
