package server

import (
	"github.com/RoundofThree/nyxeon/controllers"
	"github.com/RoundofThree/nyxeon/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)

	router.GET("/health", health.Status)
	// deal with CORS
	router.Use(middlewares.CORSMiddleware())
	// create a new unique request ID to each request
	router.Use(func(c *gin.Context) {
		uuid, _ := uuid.NewRandom()
		c.Writer.Header().Add("X-Request-Id", uuid.String())
		c.Next()
	})
	// authentication callback and validation
	oauth := router.Group("oauth")
	{
		auth := new(controllers.AuthTokenController)
		auth.Init()
		oauth.GET("/github/callback", auth.GithubOauthCallback)
		oauth.DELETE("/logout", auth.Logout)
		// oauth.POST("/verify")
		// token refresh, no need to! Right?
	}

	questGroup := router.Group("quests")
	{
		questGroup.Use(middlewares.TokenAuthMiddleware())
		quest := new(controllers.QuestController)
		quest.Init()
		questGroup.GET("/all", quest.RetrieveAll) // in the future, consider retriving pages
		questGroup.POST("/new", quest.Create)
		// questGroup.POST("/:id", quest.Update)
		questGroup.DELETE("/:id", quest.Delete)
	}
	return router
}
