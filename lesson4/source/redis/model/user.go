package model

import (
	"time"
)

// User Пользователь
type User struct {
	ID        int    `json:"id"`
	LastName  string `json:"LastName"`
	FirstName string `json:"FirstName"`
	Birth     string `json:"Birth"`
	Login     string `json:"Login"`
	Password  string `json:"Password"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (User) TableName() string {
	return "user"
}
