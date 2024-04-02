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

func FileUploadedResponse(fileId string) *fiber.Map {
	return &fiber.Map{
		"id":    fileId,
		"error": nil,
	}
}

func FileUploadFailedResponse(err error) *fiber.Map {
	return &fiber.Map{
		"id":    nil,
		"error": err.Error(),
	}
}

func GenericOKResponse() *fiber.Map {
	return &fiber.Map{
		"ok":    true,
		"error": nil,
	}
}

func GenericErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"ok":    false,
		"error": err.Error(),
	}
}
