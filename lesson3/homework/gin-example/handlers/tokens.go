package handlers

import (
	"iitu/gin-example/models"
	"iitu/gin-example/repo"
	"net/http"

	"github.com/gin-gonic/gin"
)

// getToken Получение токена
func getToken(c *gin.Context) {
	var tokenRequest models.TokenRequest
	err := c.ShouldBind(&tokenRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorJSON{
			Error: err.Error(),
		})
		return
	}
	if tokenRequest.Login == "" || tokenRequest.Password == "" {
		c.JSON(http.StatusBadRequest, models.ErrorJSON{
			Error: "login or password not passed",
		})
		return
	}

	var tokenResponse models.TokenResponse
	tokenResponse.AccessToken, err = repo.GetToken(tokenRequest.Login, tokenRequest.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorJSON{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, tokenResponse)
}
