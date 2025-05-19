package PostgresDb

import (
	"Eventify-API/internal/model"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func AuthenticateUser(db *gorm.DB, username, password string) (*model.User, error) {
	var user model.User
	result := db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			log.Printf("User not found: %s", username)
			return nil, fmt.Errorf("invalid credentials")
		}
		log.Printf("Database error: %v", result.Error)
		return nil, fmt.Errorf("error querying database: %v", result.Error)
	}

	log.Printf("Found user: %s, password_hash: %s", user.Username, user.PasswordHash)
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		log.Printf("Password mismatch for user %s: %v", username, err)
		return nil, fmt.Errorf("invalid credentials")
	}

	return &user, nil
}
