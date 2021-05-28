package middlewares

import (
	"fmt"

	"github.com/RoundofThree/nyxeon/controllers"
	"github.com/gin-gonic/gin"
)

var auth = new(controllers.AuthTokenController)

// BASED ON session token stored and mapped to user id in REDIS
func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Validating token in middleware...")
		auth.TokenValid(c)
		c.Next()
	}
}
