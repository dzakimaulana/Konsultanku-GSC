package problems

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type Handler struct {
	ProblemSvc
	Session *session.Store
}

func NewHandler(ps ProblemSvc, sess *session.Store) *Handler {
	return &Handler{
		ProblemSvc: ps,
		Session:    sess,
	}
}

func (h *Handler) AddProblem(c *fiber.Ctx) error {
	var prob AddProblem
	if err := c.BodyParser(&prob); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    fiber.StatusBadRequest,
				"message": err.Error(),
			},
		})
	}
	mseId := c.Locals("id")
	if mseId == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    fiber.StatusBadRequest,
				"message": "please login again",
			},
		})
	}
	err := h.ProblemSvc.AddProblem(c.Context(), prob, mseId.(string))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    fiber.StatusBadRequest,
				"message": err.Error(),
			},
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": fiber.Map{
			"code":    fiber.StatusOK,
			"message": "successfully create problem",
		},
	})
}

func (h *Handler) GetProblem(c *fiber.Ctx) error {
	id := c.Query("id")
	resp, err := h.ProblemSvc.GetProblem(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    fiber.StatusBadRequest,
				"message": err.Error(),
			},
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": fiber.Map{
			"code":    fiber.StatusOK,
			"message": resp,
		},
	})
}
