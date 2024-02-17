package main

import (
	"konsultanku-v2/internal/firebase/auth"
	"konsultanku-v2/internal/firebase/storage"
	"konsultanku-v2/pkg/databases"
	"konsultanku-v2/pkg/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func main() {
	ah, se, err := databases.FirebaseConn()
	if err != nil {
		panic(err.Error())
	}
	app := fiber.New()
	app.Use(logger.New())

	// Initialize session middleware
	store := session.New(session.Config{
		Expiration:        3600,
		KeyLookup:         "cookie:my-session",
		CookieSecure:      true,
		CookieSessionOnly: true,
	})

	storageRepo := storage.NewRepo(se.GetStorage())

	authRepo := auth.NewRepo(ah.GetAuth())
	authSvc := auth.NewSvc(authRepo, storageRepo)
	authHandler := auth.NewHandler(authSvc, store)
	routes.AuthRoute(authHandler, app)

	err = app.Listen(":8080")
	if err != nil {
		panic("Error starting server")
	}

}
