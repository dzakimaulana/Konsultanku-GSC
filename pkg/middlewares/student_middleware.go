package middlewares

import "github.com/gofiber/fiber/v2"

func StudentField(c *fiber.Ctx) error {
	role := c.Locals("role")
	if role != "student" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    fiber.StatusForbidden,
				"message": "This api only for student",
			},
		})
	}
	return c.Next()
}
