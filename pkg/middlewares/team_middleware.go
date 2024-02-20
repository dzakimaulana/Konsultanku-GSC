package middlewares

import "github.com/gofiber/fiber/v2"

func TeamField(c *fiber.Ctx) error {
	teamId := c.Locals("teamId")
	if teamId == nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    fiber.StatusForbidden,
				"message": "This api only for student who have a team",
			},
		})
	}
	return c.Next()
}
