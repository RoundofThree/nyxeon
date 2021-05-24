package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type QuestController struct {
}

var questModel = new(models.Quest)

func (q QuestController) RetrieveAll(c *gin.Context) {
	quests, err := questModel.GetByUserID() // by current user session
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to retrieve quests", "error": err})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{})
	return
}

func (q QuestController) Delete(c *gin.Context) {
	return
}

func (q QuestController) Create(c *gin.Context) {
	return
}
