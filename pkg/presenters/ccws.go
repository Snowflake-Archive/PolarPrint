package presenters

import "github.com/gofiber/fiber/v2"

func CcwsAccessDenied() *fiber.Map {
	return &fiber.Map{
		"ok":    false,
		"error": "authentication failed",
	}
}
