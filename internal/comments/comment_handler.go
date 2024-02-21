package comments

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type Handler struct {
	CommentSvc
	Session *session.Store
}

func NewHandler(cs CommentSvc, sess *session.Store) *Handler {
	return &Handler{
		CommentSvc: cs,
		Session:    sess,
	}
}

func (h *Handler) AddComment(c *fiber.Ctx) error {
	var comm AddComment
	if err := c.BodyParser(&comm); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    fiber.StatusBadRequest,
				"message": err.Error(),
			},
		})
	}

	studentId := c.Locals("id")
	if studentId == nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    fiber.StatusForbidden,
				"message": "please login again to get your id",
			},
		})
	}
	teamId := c.Locals("team_id")
	problemId := c.Params("problemId")
	if err := h.CommentSvc.AddComment(c.Context(), comm, studentId.(string), teamId.(string), problemId); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    fiber.StatusInternalServerError,
				"message": err.Error(),
			},
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": fiber.Map{
			"code":    fiber.StatusOK,
			"message": "successfully send comment",
		},
	})
}

func (h *Handler) GetComments(c *fiber.Ctx) error {
	mseId := c.Locals("id")
	if mseId == nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    fiber.StatusForbidden,
				"message": "please login again to get your id",
			},
		})
	}
	resp, err := h.CommentSvc.GetComments(c.Context(), mseId.(string))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    fiber.StatusInternalServerError,
				"message": err.Error(),
			},
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": fiber.Map{
			"code":    fiber.StatusOK,
			"message": resp,
		},
	})
}
