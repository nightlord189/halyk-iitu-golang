package models

// TokenRequest Запрос на получение токена
type TokenRequest struct {
	Login    string `json:"Login" form:"login"`
	Password string `json:"Password" form:"password"`
}

// TokenResponse Ответ на получение токена
type TokenResponse struct {
	AccessToken string
}
