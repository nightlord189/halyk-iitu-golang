package model

// UserState Статус пользователя
type UserState struct {
	ID          int    `json:"Id" gorm:"column:id;primary_key"`
	Name        string `json:"Name" gorm:"column:name`
	Description string `json:"Description" gorm:"column:description`
}

func (UserState) TableName() string {
	return "user_state"
}
