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

func RegisterUser(db *gorm.DB, req model.RegisterRequest) (*model.User, error) {
	var existingUser model.User
	if err := db.Where("username = ?", req.Username).First(&existingUser).Error; err == nil {
		return nil, fmt.Errorf("User already exists")
	} else if err != gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("Error checking for existing user: %v", err)

	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("Error hashing password: %v", err)
	}

	user := model.User{
		Username:     req.Username,
		PasswordHash: string(hashedPassword),
		Email:        req.Email,
		FullName:     req.FullName,
	}

	if err := db.Create(&user).Error; err != nil {
		return nil, fmt.Errorf("Error creating user: %v", err)
	}

	return &user, nil
}
