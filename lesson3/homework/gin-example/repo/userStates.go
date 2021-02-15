package repo

import (
	"errors"
	"fmt"
	"iitu/gin-example/models"
)

const userStatesFilePath string = "data/userStates.json"

var userStatesByID map[int]models.UserState = make(map[int]models.UserState)
var userStatesByName map[string]models.UserState = make(map[string]models.UserState)
var userStates []models.UserState
var userStatesMaxID int = -int(^uint(0)>>1) - 1 // min int

func uploadUserStates() error {
	err := readFile(userStatesFilePath, &userStates)
	if err != nil {
		return fmt.Errorf("upload users error error: %s", err.Error())
	}

	for _, userState := range userStates {
		userStatesByID[userState.ID] = userState
		userStatesByName[userState.Name] = userState
		if userState.ID > userStatesMaxID {
			userStatesMaxID = userState.ID
		}
	}

	return nil
}

func checkUserStatesLoaded() error {
	if userStatesByID == nil || userStatesByName == nil || userStates == nil {
		return errors.New("user states are not loaded")
	}
	return nil
}

// GetUserStateByID Получить состояние по идентификатору
func GetUserStateByID(userStateID int) (*models.UserState, error) {
	err := checkUserStatesLoaded()
	if err != nil {
		return nil, err
	}
	userState, exists := userStatesByID[userStateID]
	if !exists {
		return nil, fmt.Errorf("user state with id %d not found", userStateID)
	}

	return &userState, nil
}

// GetUserStateByName Получить состояние по наоименованию
func GetUserStateByName(userStateName string) (*models.UserState, error) {
	err := checkUserStatesLoaded()
	if err != nil {
		return nil, err
	}
	userState, exists := userStatesByName[userStateName]

	if !exists {
		return nil, nil // не найдено - это не ошибка
	}

	return &userState, nil
}

// GetUserStates Получить список всех состояний
func GetUserStates() ([]models.UserState, error) {
	err := checkUserStatesLoaded()
	if err != nil {
		return nil, err
	}

	return userStates, nil
}

// AddUserState Добавляет состояние
func AddUserState(state models.UserState) (*models.UserState, error) { // TODO: м/б сразу передавать ссылку?
	state.ID = userStatesMaxID + 1 // В любом случае инкрементится, даже если что-то указано

	if state.Name == "" {
		return nil, errors.New("user state adding failed: name not passed")
	}
	if state.Description == "" {
		return nil, errors.New("user state adding failed: description not passed")
	}

	existingState, err := GetUserStateByName(state.Name)
	if err != nil {
		return nil, err
	}
	if existingState != nil {
		return nil, fmt.Errorf("user state with name %s already exists", state.Name)
	}

	newUserStates := append(userStates, state)
	err = writeFile(userStatesFilePath, &newUserStates)
	if err != nil {
		return nil, err
	}

	userStates = newUserStates
	userStatesByID[state.ID] = state
	userStatesByName[state.Name] = state
	userStatesMaxID++

	return &state, nil
}
