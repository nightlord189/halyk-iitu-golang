package repo

import (
	"errors"
	"fmt"
	"iitu/gin-example/models"

	"github.com/google/uuid"
)

var tokens map[string]models.User

func initTokensStore() error {
	tokens = make(map[string]models.User)
	return nil
}

func createToken(user models.User) (string, error) {
	token := uuid.New().String()
	tokens[token] = user
	return token, nil
}

// GetToken Получение токена по логину/паролю
func GetToken(login, password string) (string, error) {
	if login == "" {
		return "", errors.New("login not passed")
	}
	if password == "" {
		return "", errors.New("password not passed")
	}

	usrs, err := GetUsers()
	if err != nil {
		return "", fmt.Errorf("can not get users: %v", err)
	}

	for _, user := range usrs {
		if user.Login == login && user.Password == password {
			return createToken(user)
		}
	}

	return "", fmt.Errorf("user with login %s not found or wrong password", login)
}

// CheckToken Проверка токена
func CheckToken(token string) (*models.User, error) {
	user, exists := tokens[token]
	if !exists {
		return nil, nil // error = nil токен не найден - не ошибка
	}
	return &user, nil
}
