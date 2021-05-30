package middlewares

import (
	"github.com/RoundofThree/nyxeon/config"
	"github.com/gin-gonic/gin"
)

// Middleware to handle CORS. Configure the allowed origins in config cors.origin.
func CORSMiddleware() gin.HandlerFunc {
	config := config.GetConfig()
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", config.GetString("cors.origin"))
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
