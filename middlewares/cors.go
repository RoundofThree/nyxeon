package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Setting CORS")
		c.Writer.Header().Add("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Add("Access-Control-Max-Age", "86400")
		c.Writer.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Add("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Add("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Add("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			// deal with OPTIONS
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
