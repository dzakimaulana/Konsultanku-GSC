package routes

import (
	"konsultanku-v2/internal/student"
	"konsultanku-v2/pkg/middlewares"

	"github.com/gofiber/fiber/v2"
)

func StudentRoute(h *student.Handler, f *fiber.App) {
	student := f.Group("/api/student")
	student.Use(middlewares.StudentField)
	student.Post("/profile", h.AddProfile)
	student.Post("/collaboration/:msmeId", h.AcceptOffer)
	student.Post("/comment/:problemId", h.AddComment)
	student.Post("/team/create", h.CreateTeam)
	student.Post("/team/join/:teamId", h.JoinTeam)
	student.Get("/profile", h.GetOwnProfile)
	student.Get("/problems", h.GetProblems)
	student.Get("/collaborations", h.GetCollaboration)
}
