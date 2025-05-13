package handler

import (
	"Eventify-API/internal/model"
	PostgresDb "Eventify-API/repository/db/postgres"
	"Eventify-API/service"
	"Eventify-API/utils"

	"github.com/gofiber/fiber/v2"
)
func RegisterHandler(c *fiber.Ctx) error {
    var req model.RegisterRequest 
    if err := c.BodyParser(&req); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
            Error: struct {
                Code    string `json:"code"`
                Message string `json:"message"`
            }{
                Code:    "400",
                Message: "Cannot parse request",
            },
        })
    }

    if req.Username == "" || req.Password == "" {
        return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
            Error: struct {
                Code    string `json:"code"`
                Message string `json:"message"`
            }{
                Code:    "400",
                Message: "Username and Password are required",
            },
        })
    }

    user, err:= Service.RegisterUser(req)

    if err != nil {
        if err.Error() == "User already exists" {
            return c.Status(fiber.StatusConflict).JSON(model.ErrorResponse{
                Error: struct {
                    Code    string `json:"code"`
                    Message string `json:"message"`
                }{
                    Code:    "409",
                    Message: "Username already exists",
                },
            })
        }
        return c.Status(fiber.StatusInternalServerError).JSON(model.ErrorResponse{
            Error: struct {
                Code    string `json:"code"`
                Message string `json:"message"`
            }{
                Code:    "500",
                Message: "Internal server error: " + err.Error(),
            },
        })
    }
    token, err := utils.GenerateToken(*user)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(model.ErrorResponse{
            Error: struct {
                Code    string `json:"code"`
                Message string `json:"message"`
            }{
                Code:    "500",
                Message: "Internal server error",
            },
        })
    }
    return c.Status(fiber.StatusCreated).JSON(model.RegisterResponse{
        UserId:   user.Id,
        Username: user.Username,
        Token:    token,
        Exp:      36000,
    })
    

    /*user, err := PostgresDb.RegisterUser(PostgresDb.DB, req)
    if err != nil {
        if err.Error() == "username already exists" {
            return c.Status(fiber.StatusConflict).JSON(model.ErrorResponse{
                Error: struct {
                    Code    string `json:"code"`
                    Message string `json:"message"`
                }{
                    Code:    "409",
                    Message: "Username already exists",
                },
            })
        }
        return c.Status(fiber.StatusInternalServerError).JSON(model.ErrorResponse{
            Error: struct {
                Code    string `json:"code"`
                Message string `json:"message"`
            }{
                Code:    "500",
                Message: "Internal server error: " + err.Error(),
            },
        })
    }

    return c.Status(fiber.StatusCreated).JSON(model.RegisterResponse{ // اصلاح تایپو: ResiterResponse -> RegisterResponse
        UserId:   user.Id,
        Username: user.Username,
    })*/
}
func LoginHandler(c *fiber.Ctx) error {
	var req model.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Error: struct {
				Code    string "json:\"code\""
				Message string "json:\"message\""
			}{
				Code:    "400",
				Message: "Username or Password not provided",
			},
		})
	}

	user, err := PostgresDb.AuthenticateUser(PostgresDb.DB, req.Username, req.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(model.ErrorResponse{
			Error: struct {
				Code    string "json:\"code\""
				Message string "json:\"message\""
			}{
				Code:    "401",
				Message: "Invalid credentials",
			},
		})
	}

	token, err := utils.GenerateToken(*user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.ErrorResponse{
			Error: struct {
				Code    string "json:\"code\""
				Message string "json:\"message\""
			}{
				Code:    "500",
				Message: "Internal server error",
			},
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
		"exp":   36000,
	})
}

func ReserveEventHandler(c *fiber.Ctx) error {
	println("reserve")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Event reserved successfully",
	})
}
