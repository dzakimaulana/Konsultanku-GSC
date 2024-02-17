package routes

import (
	"konsultanku-v2/internal/firebase/auth"

	"github.com/gofiber/fiber/v2"
)

func AuthRoute(h *auth.Handler, f *fiber.App) {
	auth := f.Group("/api/auth")
	auth.Post("/register", h.Register)
	auth.Post("/login", h.Login)
	auth.Post("/reset-password", h.ResetPassword)
	auth.Post("/logout", h.Logout)
}
