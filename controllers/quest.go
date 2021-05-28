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

func (q QuestController) RetrieveAll(c *gin.Context) {
	fmt.Println("I am here")
	userID := c.Keys["userID"]
	user, err := q.userManager.GetByUserID(userID.(string))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid session"})
		return
	}
	quests, err := q.questManager.GetByUser(user) // by current user session
	fmt.Println("Retrieved quests: ", quests)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Error to retrieve quests", "error": err})
		return
	}
	// pass quests as JSON
	c.JSON(http.StatusOK, quests)
	return
}

func (q QuestController) Delete(c *gin.Context) {
	return
}

// TODO
func (q QuestController) Create(c *gin.Context) {
	return
}
