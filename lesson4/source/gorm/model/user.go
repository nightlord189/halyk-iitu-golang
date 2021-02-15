package model

import (
	"time"
)

// User Пользователь
type User struct {
	ID        int    `json:"id" gorm:"column:id;primary_key"`
	LastName  string `json:"LastName" gorm:"column:lastname;index"`
	FirstName string `json:"FirstName" gorm:"column:firstname"`
	Birth     string `json:"Birth" gorm:"column:birth;"`
	Login     string `json:"Login" gorm:"column:login;"`
	Password  string `json:"Password" gorm:"column:password"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (User) TableName() string {
	return "user"
}
