package model

import "time"

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type User struct {
	Id           int64  `gorm:"primaryKey" json:"id"`
	Email        string `gorm:"unique" json:"email"`
	Username     string `gorm:"unique" json:"username"`
	PasswordHash string `json:"-"`
	FullName     string `json:"full_name"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email,omitempty"`
	FullName string `json:"fullname,omitempty"`
}
type RegisterResponse struct {
	UserId   int64  `json:"user_id"`
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
