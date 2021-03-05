package handler

import (
	"iitu/swagger-example/model"

	"github.com/gin-gonic/gin"
)

// GetUser godoc
// @Summary Get user
// @Description get sample user from in-memory logic
// @Tags Users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Param email query string false "search by email" Format(email)
// @Success 200 {object} model.User
// @Failure 404 {object} model.Response
// @Router /user/{id} [get]
func GetUser(c *gin.Context) {
	id := c.Param("id")
	if id == "159" {
		user := model.User{
			ID:        13218,
			Name:      "User1",
			Lastname:  "Ivanov",
			Email:     "user@mail.com",
			IsAdmin:   false,
			AuthCount: 159,
		}
		c.JSON(200, user)
		return
	} else {
		c.JSON(404, model.Response{
			Code:    -1,
			Message: "not found",
		})
	}
}

// UpdateUser godoc
// @Summary Update user
// @Description update user in example servuce
// @Tags Users
// @Accept  json
// @Produce  json
/// @Param user body model.User true "Model to update"
// @Success 200 {object} model.Response
// @Failure 400 {object} model.Response
// @Router /user [post]
func UpdateUser(c *gin.Context) {
	c.JSON(200, model.Response{
		Code:    0,
		Message: "success",
	})
}
