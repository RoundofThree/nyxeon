package controllers

import (
	"net/http"

	"github.com/RoundofThree/nyxeon/models"
	"github.com/gin-gonic/gin"
)

type AuthTokenController struct {
}

var sessionManager = new(models.SessionManager)

// Validate session token sent by the client and restore session in the request.
// This is injected as middleware.
func (ctl AuthTokenController) TokenValid(c *gin.Context) {
	// extract token from cookie
	cookie, err := c.Request.Cookie("nyx_sess_id")
	if err != nil {
		if err == http.ErrNoCookie {
			c.Writer.WriteHeader(http.StatusUnauthorized)
			return
		}
		c.Writer.WriteHeader(http.StatusBadRequest)
	}
	token := cookie.Value
	// check the session token in Redis
	userID, err := sessionManager.FetchSession(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
		return
	}
	// access in other handlers in c.Keys
	c.Set("userID", userID)
}

// Deletes the session in server cache.
func (ctl AuthTokenController) Logout(c *gin.Context) {
	// Delete session in redis
}

func (ctl AuthTokenController) GithubOauthCallback(c *gin.Context) {
	// get the code

	// request Github API

	// store session to Redis

	// set cookie nyx_sess_id
	// send HTTP Found to client side dashboard url
}
