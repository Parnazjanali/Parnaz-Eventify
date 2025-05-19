package Service

import (
	"Eventify-API/internal/model"
	PostgresDb "Eventify-API/repository/db/postgres"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2/log"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(req model.RegisterRequest) (*model.User, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, fmt.Errorf("error creating logger: %v", err)
	}
	defer logger.Sync()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.Error("Failed to hash password", zap.Error(err))
		return nil, fmt.Errorf("failed to hash password: %v", err)
	}

	user := model.User{
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
		FullName:     req.FullName,
	}
	if err := PostgresDb.DB.Create(&user).Error; err != nil {
		logger.Error("Failed to create user", zap.Error(err))
		return nil, fmt.Errorf("failed to create user: %v", err)
	}
	logger.Info("User created", zap.Int64("user_id", user.Id))

	profilePayload := map[string]interface{}{
		"userId":   user.Id,
		"email":    user.Email,
		"username": user.Username,
		"fullName": req.FullName,
	}

	body, _ := json.Marshal(profilePayload)

	resp, err := http.Post("http://127.0.0.1:8083/profiles/", "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Error("Failed to create profile", zap.Error(err))
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	logger.Info("Profile manager response", zap.String("body", string(bodyBytes)))

	if resp.StatusCode != http.StatusCreated {
		logger.Error("Profile creation failed", zap.Int("status_code", resp.StatusCode))
		return &user, fmt.Errorf("profile creation failed")
	}

	logger.Info("Profile created successfully")
	return &user, nil
}

func AuthenticateUser(username, password string)(*model.User, error){
	logger, err:= zap.NewProduction()
	if err!= nil{
		return nil, fmt.Errorf("error creating logger: %v", err)
	}
	defer logger.Sync()

	var req model.User
	if err := PostgresDb.DB.Where("username = ?", username).First(&req).Error; err != nil {
		logger.Error("User not found", zap.Error(err))
		return nil, fmt.Errorf("user not found")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(req.PasswordHash), []byte(password)); err != nil {
		logger.Error("Invalid password", zap.Error(err))
		return nil, fmt.Errorf("invalid password")
	}

	reqPayLoad := map[string]interface{}{
		"username": username,
		"password": password,
	}

	body, _:= json.Marshal(reqPayLoad)

	resp , err:= http.Post("http://127.0.0.1:8083/profiles/validate", "application/json", bytes.NewBuffer(body))
	if err!= nil{
		log.Error("Failed to validate profile", zap.Error(err))
	}
	defer resp.Body.Close()

	bodyBytes, _:= io.ReadAll(resp.Body)
	logger.Info("Profile manager response", zap.String("body", string(bodyBytes)))
	if resp.StatusCode != http.StatusOK {
		logger.Error("Profile validation failed", zap.Int("status_code", resp.StatusCode))
		return nil, fmt.Errorf("profile validation failed")
	}
	logger.Info("Profile validated successfully")
	
	return &req, nil
}
