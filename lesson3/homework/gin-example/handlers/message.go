package handlers

import (
	"github.com/gin-gonic/gin"
	"iitu/gin-example/models"
	"net/http"
)

func processMessage (c *gin.Context) {
	var message models.UserMessage
	err := c.Bind(&message)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorJSON{
			Error: err.Error(),
		})
		return
	}

	if message.Message == "" {
		c.JSON(http.StatusOK, "You are silent!")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"You say": message.Message,
	})
}
