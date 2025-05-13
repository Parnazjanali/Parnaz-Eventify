package model

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type User struct {
	Id           int    `json:"id" gorm:"primaryKey"`
    Username     string `json:"username" gorm:"unique"`
    PasswordHash string `json:"password_hash"`
    Email        string `json:"email,omitempty"`
    FullName     string `json:"fullName,omitempty"`
}
type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email   string `json:"email,omitempty"`
	FullName string `json:"fullName",omitempty`
}
type RegisterResponse struct {
	UserId int    `json:"user_id"`
	Username string `json:"username"`
	Token    string `json:"token"`
	Exp      int    `json:"exp"`
}
type ErrorResponse struct {
	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}
