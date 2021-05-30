package middlewares

import (
	"github.com/RoundofThree/nyxeon/controllers"
	"github.com/gin-gonic/gin"
)

var auth = new(controllers.AuthTokenController)

// Middleware to check if the session token is valid.
func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth.TokenValid(c)
		c.Next()
	}
}
