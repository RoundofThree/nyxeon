package config

import (
	"log"
	"path/filepath"

	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var config *viper.Viper
var oauthCfg *oauth2.Config

func Init(env string) {
	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName(env)
	config.AddConfigPath("../config/")
	config.AddConfigPath("config/")
	err := config.ReadInConfig()
	if err != nil {
		log.Fatal("Error parsing config files", err)
	}

	oauthCfg = &oauth2.Config{
		ClientID:     config.GetString("oauth.githubClientID"),
		ClientSecret: config.GetString("oauth.githubClientSecret"),
		Endpoint:     github.Endpoint,
		Scopes:       config.GetStringSlice("oauth.githubScopes"),
	}
}

func relativePath(basedir string, path *string) {
	p := *path
	if len(p) > 0 && p[0] != '/' {
		*path = filepath.Join(basedir, p)
	}
}

func GetConfig() *viper.Viper {
	return config
}

func GetOauthConfig() *oauth2.Config {
	return oauthCfg
}
