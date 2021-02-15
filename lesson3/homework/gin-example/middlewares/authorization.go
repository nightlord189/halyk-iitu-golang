package middlewares

import (
	"errors"
	"github.com/gin-gonic/gin"
	"iitu/gin-example/repo"
	"net/http"
	"strings"
)

func CheckToken(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		c.AbortWithError(http.StatusBadRequest, errors.New("bad authorization schema, need Bearer"))
		return
	}

	if len(authHeader) < 9 {
		c.AbortWithError(http.StatusUnauthorized, errors.New("token not passed"))
		return
	}

	token := authHeader[7:]

	user, err := repo.CheckToken(token)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if user == nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Set("user", *user)
}
