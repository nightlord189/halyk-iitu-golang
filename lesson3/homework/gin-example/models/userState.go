package models

// UserState Статус пользователя
type UserState struct {
	ID          int    `json:"Id"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
}