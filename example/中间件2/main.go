package main

import (
	"log"
	"time"
)

import (
	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		// Set example varible
		c.Set("example", "12345")
		// before request
		c.Next()
		// after request
		latency := time.Since(t)
		log.Print(latency)
		//access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}

}

func main() {
	r := gin.New()
	r.Use(Logger())

	r.GET("/test", func(c *gin.Context) {
		example := c.MustGet("example").(string)
		log.Println(example + "================")

	})
	r.Run(":3333")

}

/*
main函数中路由r使用了logger函数，所以在路由执行的时候这个函数肯定会起作用的，我们再来看看logger函数

logger函数中主要干了以下几件事

1）request执行前，给上下文赋值example

2）使用next，使继续

3）request执行后，打印请求执行时间，用的是go的Since函数

4）打印访问状态
*/
