package repo

import (
	"errors"
	"fmt"
	"iitu/gin-example/models"
)

var users []models.User
var usersLen int

func uploadUsers() error {
	//users = make([]models.User, 10)
	err := readFile("data/users.json", &users)
	if err != nil {
		return fmt.Errorf("upload users error: %s", err.Error())
	}
	usersLen = len(users)

	for i := 0; i < usersLen; i++ {
		var state *models.UserState
		state, err = GetUserStateByID(users[i].StateID)
		if err != nil {
			return fmt.Errorf("upload users error: %s", err.Error())
		}
		users[i].State = state
	}
	return nil
}

func checkUsersLoaded() error {
	if users == nil {
		return errors.New("users are not loaded")
	}
	return nil
}

// GetUsers Получить список юзеров
func GetUsers() ([]models.User, error) {
	err := checkUsersLoaded()
	return users, err
}

// GetUserByID Получить юзера по его идентификатору
func GetUserByID(userID int) (*models.User, error) {
	if err := checkUsersLoaded(); err != nil {
		return nil, err
	}

	for _, user := range users {
		if user.ID == userID {
			return &user, nil
		}
	}

	return nil, fmt.Errorf("user with id %d not found", userID)
}
