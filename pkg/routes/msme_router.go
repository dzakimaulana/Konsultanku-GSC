package routes

import (
	"konsultanku-v2/internal/msme"
	"konsultanku-v2/pkg/middlewares"

	"github.com/gofiber/fiber/v2"
)

func MsmeRoute(h *msme.Handler, f *fiber.App) {
	msme := f.Group("/api/msme")
	msme.Use(middlewares.MsmeField)
	msme.Post("/profile", h.AddProfile)
	msme.Post("/problem", h.AddProblem)
	msme.Post("/collaboration/:studentId", h.AddedCollab)
	msme.Put("/progress/:studentId", h.GiveProgress)
	msme.Put("/end-collaboration/:studentId", h.EndCollaboration)
	msme.Get("/profile", h.GetOwnProfile)
	msme.Get("/comments", h.GetComments)
}
