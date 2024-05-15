package middleware

import (
	"github.com/DangNguyen146/bloggolangreactjs/tree/main/blogbe/util"
	"github.com/gofiber/fiber/v2"
)

func IsAuthorized(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	if _, err := util.ParseJwt(cookie); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	return c.Next()
}
