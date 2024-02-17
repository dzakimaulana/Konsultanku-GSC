package middlewares

import "github.com/gofiber/fiber/v2"

func MsmeField(c *fiber.Ctx) error {
	role := c.Locals("role")
	if role != "msme" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    fiber.StatusForbidden,
				"message": "This api only for msme",
			},
		})
	}
	return c.Next()
}
