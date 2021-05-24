package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

var github_config *oauth2.Config

func GithubOauthLogin(c *gin.Context) {
	conf, state := getGithubOauthURL()
	redirectURL := conf.AuthCodeURL(state)

	session := sessions.Default(c)
	session.Set("state", state)
	err := session.Save()
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Redirect(http.StatusSeeOther, redirectURL)
}

// callback
