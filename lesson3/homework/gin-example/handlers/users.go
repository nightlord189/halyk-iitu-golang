package handlers

import (
	"iitu/gin-example/models"
	"iitu/gin-example/repo"
	"net/http"

	"github.com/gin-gonic/gin"
)

// getUsers Получить список всех пользователей
func getUsers(c *gin.Context) {
	users, err := repo.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorJSON{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, users)
}
