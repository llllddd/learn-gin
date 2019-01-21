package main

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func respondWithError(code int, message string, c *gin.Context) {
	resp := map[string]string{"error": message}
	c.JSON(code, resp)
	c.Abort()
}

func TokenAuthHiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.FormValue("api_token")
		if token == "" {
			respondWithError(401, "API token required", c)
			return
		}
		if token != os.Getenv("TEMP") {
			respondWithError(401, "invalid API token", c)
			return
		}
		c.String(http.StatusOK, "hello, world")
		c.Writer.Header().Set("X-Request-ID", string(time.Second)) //通过这种方式给resp头放一些东西，比如一些校验类
		c.Next()
	}
}

func main() {
	r := gin.New()
	r.Use(TokenAuthHiddleware())
	r.Run(":3333")
}
