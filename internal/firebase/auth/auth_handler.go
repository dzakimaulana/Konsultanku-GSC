package auth

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type Handler struct {
	AuthSvc
	Session *session.Store
}

func NewHandler(sa AuthSvc, sess *session.Store) *Handler {
	return &Handler{
		AuthSvc: sa,
		Session: sess,
	}
}

func (h *Handler) Register(c *fiber.Ctx) error {
	// Parse request body into RegisterFirebaseReq struct
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    fiber.StatusBadRequest,
				"message": err.Error(),
			},
		})
	}

	var user RegisterFirebaseReq
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    fiber.StatusBadRequest,
				"message": err.Error(),
			},
		})
	}
	user.File = file

	res, err := h.AuthSvc.Register(c.Context(), user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    fiber.StatusBadRequest,
				"message": err.Error(),
			},
		})
	}

	// Return a success response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": fiber.Map{
			"code": fiber.StatusOK,
			"data": fmt.Sprintf("Email verification already send into your email: %s", res.Email),
		},
	})
}

func (h *Handler) Login(c *fiber.Ctx) error {
	var log LoginReq
	if err := c.BodyParser(&log); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    fiber.StatusBadRequest,
				"message": err.Error(),
			},
		})
	}

	resp, err := h.AuthSvc.Login(c.Context(), log)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    fiber.StatusBadRequest,
				"message": err.Error(),
			},
		})
	}

	// set session
	sess, err := h.Session.Get(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    fiber.StatusBadRequest,
				"message": err.Error(),
			},
		})
	}
	sess.Set("access_token", resp.IDToken)
	sess.Set("refresh_token", resp.RefreshToken)
	sess.Set("uid", resp.LocalID)
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
			"code": fiber.StatusOK,
			"data": resp,
		},
	})
}

func (h *Handler) ResetPassword(c *fiber.Ctx) error {
	var email RPReq
	if err := c.BodyParser(&email); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    fiber.StatusBadRequest,
				"message": err.Error(),
			},
		})
	}

	res, err := h.AuthSvc.ResetPassword(c.Context(), email)
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
			"code": fiber.StatusOK,
			"data": fmt.Sprintf("Check inbox in your email to reset password: %s", res.Email),
		},
	})
}

func (h *Handler) Logout(c *fiber.Ctx) error {
	sess, err := h.Session.Get(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    fiber.StatusBadRequest,
				"message": err.Error(),
			},
		})
	}

	// revoke refresh token from firebase
	uid := sess.Get("uid").(string)
	if err := h.AuthSvc.Logout(c.Context(), uid); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    fiber.StatusBadRequest,
				"message": err.Error(),
			},
		})
	}
	// destroy session in server
	sess.Destroy()
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": fiber.Map{
			"code": fiber.StatusOK,
			"data": "Successfully logout.",
		},
	})
}
