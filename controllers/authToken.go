package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/RoundofThree/nyxeon/config"
	"github.com/RoundofThree/nyxeon/models"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/github"
	"github.com/google/uuid"
)

type AuthTokenController struct {
	sessionManager *models.SessionManager
	userManager    *models.UserManager
}

func (ctl AuthTokenController) Init() {
	ctl.sessionManager = new(models.SessionManager)
	ctl.userManager = new(models.UserManager)
}

func getSessionFromCookie(r *http.Request) (string, error) {
	cookie, err := r.Cookie("nyx_sess_id")
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}

// Validate session token sent by the client and restore session in the request.
// This is injected as middleware.
func (ctl AuthTokenController) TokenValid(c *gin.Context) {
	// extract token from cookie
	token, err := getSessionFromCookie(c.Request)
	if err != nil {
		if err == http.ErrNoCookie {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
			return
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{})
		return
	}
	// check the session token in Redis
	userID, err := ctl.sessionManager.FetchSession(token)
	fmt.Println("UserID from session: ", userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
		return
	}
	// access in other handlers in c.Keys
	c.Set("userID", userID)
}

// Deletes the session in server cache.
func (ctl AuthTokenController) Logout(c *gin.Context) {
	// extract token from cookie
	token, err := getSessionFromCookie(c.Request)
	if err != nil {
		if err == http.ErrNoCookie {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
			return
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{})
		return
	}
	// Delete session in redis
	ctl.sessionManager.DeleteSession(token)
	// make client delete session cookie
	c.SetCookie("nyx_sess_id", "", 0, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "session deleted"})
}

// to defend against possible CSRF, attach server generated state to the callback URL
/*
func (ctl AuthTokenController) StartOauth(c *gin.Context) {
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)

}
*/

func (ctl AuthTokenController) GithubOauthCallback(c *gin.Context) {
	// get the code
	err := c.Request.ParseForm()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Could not parse query"})
		return
	}
	code := c.Request.FormValue("code")
	fmt.Println("Code is ", code)
	// request Github API
	token, err := config.GetOauthConfig().Exchange(context.Background(), code)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Could not retrieve token", "error": err})
		return
	}
	if !token.Valid() {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid token"})
		return
	}
	client := github.NewClient(config.GetOauthConfig().Client(context.Background(), token))
	user, _, err := client.Users.Get(context.Background(), "")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "cannot retrieve userid"})
		return
	}
	// store session to Redis
	newUUID, err := uuid.NewRandom()
	sessionToken := newUUID.String()
	fmt.Println("Token is ", sessionToken)
	fmt.Println("User is ", *user)
	fmt.Println("Email is ", user.GetEmail())
	err = ctl.sessionManager.UpdateSession(sessionToken, user.GetEmail())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{})
		return
	}
	// db create user if not present
	_, err = ctl.userManager.GetByUserID(user.GetEmail())
	if err != nil {
		ctl.userManager.CreateUser(user.GetEmail())
	}
	// set cookie nyx_sess_id
	c.SetCookie("nyx_sess_id", sessionToken, 60*60*24, "/", "localhost", false, true)
	// send HTTP Found to client side dashboard url
	c.Redirect(http.StatusFound, "http://localhost:3000/dashboard")
}
