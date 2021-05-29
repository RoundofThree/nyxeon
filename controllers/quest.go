package controllers

import (
	"fmt"
	"net/http"

	"github.com/RoundofThree/nyxeon/models"
	"github.com/gin-gonic/gin"
)

type QuestController struct {
	userManager  *models.UserManager
	questManager *models.QuestManager
}

func (ctl QuestController) Init() {
	ctl.userManager = new(models.UserManager)
	ctl.questManager = new(models.QuestManager)
}

func (ctl QuestController) RetrieveAll(c *gin.Context) {
	fmt.Println("I am here")
	userID := c.Keys["userID"]
	user, err := ctl.userManager.GetByUserID(userID.(string))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid session"})
		return
	}
	quests, err := ctl.questManager.GetByUser(user) // by current user session
	fmt.Println("Retrieved quests: ", quests)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Error to retrieve quests", "error": err})
		return
	}
	// pass quests as JSON
	c.JSON(http.StatusOK, quests)
	return
}

// TODO
func (q QuestController) Delete(c *gin.Context) {
	return
}

// Create a quest given the content and the categories. The content is a large text, categories
// is an array of tags.
func (ctl QuestController) Create(c *gin.Context) {
	userID := c.Keys["userID"]
	user, err := ctl.userManager.GetByUserID(userID.(string))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid session"})
		return
	}
	// form content structure: {categories: [string], content: string}
	err = c.Request.ParseMultipartForm(1000)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{})
	}
	content := c.Request.Form.Get("content")
	categories := c.Request.Form["categories[]"]
	if err = ctl.questManager.Create(user, content, categories); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{})
	}
	c.JSON(http.StatusCreated, gin.H{})
}
