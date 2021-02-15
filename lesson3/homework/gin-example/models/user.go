package models

// User Пользователь
type User struct {
	ID        int        `json:"id"`
	LastName  string     `json:"LastName"`
	FirstName string     `json:"FirstName"`
	Birth     string     `json:"Birth"`
	Login     string     `json:"Login"`
	Password  string     `json:"Password"`
	StateID   int        `json:"StateID"`
	State     *UserState `json:"State"`
}
