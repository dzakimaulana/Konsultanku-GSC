package msme

import (
	"konsultanku-v2/internal/comments"
	"konsultanku-v2/internal/problems"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type Handler struct {
	MsmeSvc
	ProblemSvc problems.ProblemSvc
	CommentSvc comments.CommentSvc
	Session    *session.Store
}

func NewHandler(ms MsmeSvc, ps problems.ProblemSvc, cs comments.CommentSvc, sess *session.Store) *Handler {
	return &Handler{
		MsmeSvc:    ms,
		ProblemSvc: ps,
		Session:    sess,
	}
}

func (h *Handler) GetOwnProfile(c *fiber.Ctx) error {
	id := c.Locals("id")
	resp, err := h.MsmeSvc.GetOwnProfile(c.Context(), id.(string))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    fiber.StatusBadRequest,
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

func (h *Handler) AddProfile(c *fiber.Ctx) error {
	var addReq AddReq
	if err := c.BodyParser(&addReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    fiber.StatusBadRequest,
				"message": err.Error(),
			},
		})
	}
	id := c.Locals("id")
	if err := h.MsmeSvc.AddProfile(c.Context(), addReq, id.(string)); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    fiber.StatusBadRequest,
				"message": err.Error(),
			},
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": fiber.Map{
			"code":    fiber.StatusOK,
			"message": "successfully added msme profile",
		},
	})
}

func (h *Handler) AddedCollab(c *fiber.Ctx) error {
	studentId := c.Params("studentId")
	if studentId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    fiber.StatusBadRequest,
				"message": "you need student's ID for collaboration",
			},
		})
	}
	msmeId := c.Locals("id").(string)
	if err := h.MsmeSvc.AddedCollab(c.Context(), studentId, msmeId); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    fiber.StatusBadRequest,
				"message": err.Error(),
			},
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": fiber.Map{
			"code":    fiber.StatusOK,
			"message": "successfully send offer to student",
		},
	})
}

func (h *Handler) AddProblem(c *fiber.Ctx) error {
	var prob problems.AddProblem
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

func (h *Handler) GiveProgress(c *fiber.Ctx) error {
	var progress UpdateProgress
	if err := c.BodyParser(&progress); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    fiber.StatusBadRequest,
				"message": err.Error(),
			},
		})
	}
	msmeId := c.Locals("id")
	if msmeId == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    fiber.StatusBadRequest,
				"message": "please login again to get your id",
			},
		})
	}
	studentId := c.Params("studentId")
	if studentId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    fiber.StatusBadRequest,
				"message": "student id not detected",
			},
		})
	}
	if err := h.MsmeSvc.GiveProgress(c.Context(), progress, msmeId.(string), studentId); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    fiber.StatusBadRequest,
				"message": err.Error(),
			},
		})
	}
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"success": fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "successfully give student progress",
		},
	})
}

func (h *Handler) EndCollaboration(c *fiber.Ctx) error {
	var end EndCollaboration
	if err := c.BodyParser(&end); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    fiber.StatusBadRequest,
				"message": err.Error(),
			},
		})
	}
	msmeId := c.Locals("id")
	if msmeId == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    fiber.StatusBadRequest,
				"message": "please login again to get your id",
			},
		})
	}
	studentId := c.Params("studentId")
	if studentId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    fiber.StatusBadRequest,
				"message": "student id not detected",
			},
		})
	}
	if err := h.MsmeSvc.EndCollaboration(c.Context(), end, msmeId.(string), studentId); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    fiber.StatusBadRequest,
				"message": err.Error(),
			},
		})
	}
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"success": fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "successfully end collaboration",
		},
	})
}
