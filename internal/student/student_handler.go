package student

import (
	"konsultanku-v2/internal/comments"
	"konsultanku-v2/internal/problems"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type Handler struct {
	StudentSvc
	ProblemSvc problems.ProblemSvc
	CommentSvc comments.CommentSvc
	Session    *session.Store
}

func NewHandler(ss StudentSvc, ps problems.ProblemSvc, cs comments.CommentSvc, sess *session.Store) *Handler {
	return &Handler{
		StudentSvc: ss,
		ProblemSvc: ps,
		Session:    sess,
	}
}

func (h *Handler) AddProfile(c *fiber.Ctx) error {
	var addReq AddStudent
	if err := c.BodyParser(&addReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    fiber.StatusBadRequest,
				"message": err.Error(),
			},
		})
	}
	id := c.Locals("id")
	if id == nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    fiber.StatusForbidden,
				"message": "please login again, your id not found",
			},
		})
	}
	if err := h.StudentSvc.AddProfile(c.Context(), addReq, id.(string)); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    fiber.StatusForbidden,
				"message": "please login again, your id not found",
			},
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": fiber.Map{
			"code":    fiber.StatusOK,
			"message": "successfulyy create student profile",
		},
	})
}

func (h *Handler) GetOwnProfile(c *fiber.Ctx) error {
	id := c.Locals("id")
	if id == nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    fiber.StatusForbidden,
				"message": "please login again, your id not found",
			},
		})
	}
	resp, err := h.StudentSvc.GetOwnProfile(c.Context(), id.(string))
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

func (h *Handler) AcceptOffer(c *fiber.Ctx) error {
	mseId := c.Params("mseId")
	if mseId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    fiber.StatusBadRequest,
				"message": "you need mse's ID for collaboration",
			},
		})
	}
	studentId := c.Locals("id")
	if studentId == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    fiber.StatusBadRequest,
				"message": "please login again to get your id",
			},
		})
	}
	phoneNumber, err := h.StudentSvc.AcceptOffer(c.Context(), mseId, studentId.(string))
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
			"message": phoneNumber,
		},
	})
}

func (h *Handler) GetProblems(c *fiber.Ctx) error {
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

func (h *Handler) AddComment(c *fiber.Ctx) error {
	var comm comments.AddComment
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
	if err := h.CommentSvc.AddComment(c.Context(), comm, studentId.(string), teamId.(string)); err != nil {
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
			"message": "successfully send comment",
		},
	})
}

func (h *Handler) GetCollaboration(c *fiber.Ctx) error {
	studentId := c.Locals("id")
	if studentId == nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    fiber.StatusForbidden,
				"message": "please login again to get your id",
			},
		})
	}
	resp, err := h.StudentSvc.GetCollaboration(c.Context(), studentId.(string))
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

func (h *Handler) CreateTeam(c *fiber.Ctx) error {
	var team CreateTeam
	if err := c.BodyParser(&team); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    fiber.StatusBadRequest,
				"message": err.Error(),
			},
		})
	}
	teamResp, err := h.StudentSvc.CreateTeam(c.Context(), team)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    fiber.StatusInternalServerError,
				"message": err.Error(),
			},
		})
	}
	// set teamId into session
	sess, err := h.Session.Get(c)
	if err != nil {
		return err
	}
	sess.Set("teamId", teamResp.ID)
	if err := sess.Save(); err != nil {
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
			"message": "successfully create a team",
		},
	})
}

func (h *Handler) JoinTeam(c *fiber.Ctx) error {
	studentId := c.Locals("id")
	if studentId == nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    fiber.StatusForbidden,
				"message": "please login again to get your id",
			},
		})
	}
	teamId := c.Params("teamId")
	if teamId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    fiber.StatusBadRequest,
				"message": "you need team's ID for collaboration",
			},
		})
	}
	if err := h.StudentSvc.JoinTeam(c.Context(), teamId, studentId.(string)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    fiber.StatusInternalServerError,
				"message": err.Error(),
			},
		})
	}
	// set teamId into session
	sess, err := h.Session.Get(c)
	if err != nil {
		return err
	}
	sess.Set("teamId", teamId)
	if err := sess.Save(); err != nil {
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
			"message": "successfully join team",
		},
	})
}
