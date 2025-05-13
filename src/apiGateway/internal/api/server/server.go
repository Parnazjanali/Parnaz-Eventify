package Server

import (
	"Eventify-API/internal/api/handler"
	"Eventify-API/internal/api/middleware"
	"log"

	"github.com/gofiber/fiber/v2"
)

func SetupApi() {
	app := fiber.New()

    registerGroup := app.Group("/register")
    registerGroup.Post("/user", handler.RegisterHandler)

	authGroup := app.Group("/auth")
	authGroup.Post("/login", handler.LoginHandler)

	reserveGroup := app.Group("/reserve")
	reserveGroup.Post("/event", middleware.AuthMiddleware, handler.ReserveEventHandler)

	port := "127.0.0.1:8080"
	log.Fatal(app.Listen(port))
}
