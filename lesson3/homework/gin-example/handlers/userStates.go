package handlers

import (
	"fmt"
	"iitu/gin-example/models"
	"iitu/gin-example/repo"
	"net/http"

	"github.com/gin-gonic/gin"
)

// addUserState Добавляет новое состояние
func addUserState(c *gin.Context) {
	var userState models.UserState
	err := c.ShouldBind(&userState)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorJSON{
			Error: fmt.Sprintf("bad user state format: %v", err),
		})
		return
	}

	addedUserState, err := repo.AddUserState(userState)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorJSON{
			Error: fmt.Sprintf("bad user state format: %v", err),
		})
		return
	}
	if addedUserState == nil {
		c.JSON(http.StatusInternalServerError, models.ErrorJSON{
			Error: "somethig wrong happend, user state don't added",
		})
		return
	}

	c.JSON(http.StatusCreated, addedUserState)
}
