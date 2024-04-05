package presenters

import "github.com/gofiber/fiber/v2"

func MissingAuthorization() *fiber.Map {
	return &fiber.Map{
		"error": "Missing Authorization",
	}
}

func Unauthorized() *fiber.Map {
	return &fiber.Map{
		"error": "Unauthorized",
	}
}
