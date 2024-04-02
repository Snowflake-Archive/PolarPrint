package presenters

import (
	"github.com/gofiber/fiber/v2"
)

func FileListResponse(data []string) *fiber.Map {
	return &fiber.Map{
		"files": data,
		"error": nil,
	}
}

func FileListErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"files": nil,
		"error": err.Error(),
	}
}
