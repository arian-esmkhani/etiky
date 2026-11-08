package middleware

import (
	"fiber-app/helper"

	"github.com/gofiber/fiber/v2"
)

func Auth(jwt *helper.JWT) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get token from cookie
		token := c.Cookies("access_token")
		
		if token == "" {
			return c.Status(401).JSON(fiber.Map{
				"error": "Please log in",
				"code":  "NO_TOKEN",
			})
		}

		// Validate token
		claims, err := jwt.ParseAccessToken(token)
		if err != nil {
			// Remove invalid cookie
			c.ClearCookie("access_token")
			return c.Status(401).JSON(fiber.Map{
				"error": "Session expired, please log in again",
				"code":  "INVALID_TOKEN",
			})
		}

		// Store user info in context
		c.Locals("userId", claims.UserId)
		c.Locals("role", claims.Role)

		return c.Next()
	}
}