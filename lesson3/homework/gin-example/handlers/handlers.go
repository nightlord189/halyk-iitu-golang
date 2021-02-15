package handlers

import (
	"errors"
	"iitu/gin-example/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

// MapRoutes Создает маршруты
func MapRoutes(router *gin.Engine) error {
	if router == nil {
		return errors.New("o my god, router not created")
	}

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	router.POST("/message", processMessage)

	tokensRouterGroup := router.Group("tokens")
	{
		tokensRouterGroup.POST("/", getToken)
	}

	usersRouterGroup := router.Group("users")
	{
		usersRouterGroup.Use(middlewares.CheckToken)
		usersRouterGroup.GET("/", getUsers)
	}

	userStateRouterGroup := router.Group("user-states")
	{
		userStateRouterGroup.Use(middlewares.CheckToken)
		userStateRouterGroup.POST("/", addUserState)
	}

	return nil
}
