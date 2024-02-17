package middlewares

import (
	"konsultanku-v2/pkg/databases"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func GetSession(sess *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sess, err := sess.Get(c)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": fiber.Map{
					"code":    fiber.StatusUnauthorized,
					"message": err.Error(),
				},
			})
		}
		token := sess.Get("access_token")
		if token == "" || token == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": fiber.Map{
					"code":    fiber.StatusUnauthorized,
					"message": err.Error(),
				},
			})
		}
		_, err = databases.AuthMd.VerifyIDToken(c.Context(), token.(string))
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": fiber.Map{
					"code":    fiber.StatusUnauthorized,
					"message": err.Error(),
				},
			})
		}
		uid := sess.Get("uid")
		userinfo, err := databases.AuthMd.GetUser(c.Context(), uid.(string))
		role := userinfo.CustomClaims["role"].(string)
		c.Locals("role", role)
		return c.Next()
	}
}
