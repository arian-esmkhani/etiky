package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func RequireAdmin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userRole := c.Locals("role")
		
		if userRole != "admin" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message": "Admin access required",
			})
		}
		
		return c.Next()
	}
}