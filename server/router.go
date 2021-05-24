package server

import (
	"github.com/RoundofThree/nyxeon/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)

	router.GET("/health", health.Status)
	// deal with session with middlewares
	router.Use(middlewares.AuthMiddleware())

	oauth := router.Group("oauth")
	{
		oauth.GET("/github", utils.GithubOauthLogin)
	}

	callback := router.Group("callback")
	{
		callback.GET("/github", utils.GithubCallback)
	}

	questGroup := router.Group("quests")
	{
		quest := new(controllers.QuestController)
		questGroup.GET("/", quest.RetrieveAll) // in the future, consider retriving pages
		questGroup.POST("/", quest.Create)
		// questGroup.POST("/:id", quest.Update)
		questGroup.DELETE("/:id", quest.Delete)
	}
	return router
}
