package presenters

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func CcwsAccessDenied() *fiber.Map {
	return &fiber.Map{
		"ok":    false,
		"error": "authentication failed",
	}
}

func CcwsUnknownError(err error) *fiber.Map {
	return &fiber.Map{
		"event": "error",
		"error": err.Error(),
	}
}

func CcwsUnrecognizedPacket(requestedType string) *fiber.Map {
	return &fiber.Map{
		"event": "error",
		"error": fmt.Sprintf("unrecognized packet type %s", requestedType),
	}
}

func CcwsSyncPrintersFailure(err error) *fiber.Map {
	return &fiber.Map{
		"event": "sync_printers_finish",
		"error": err.Error(),
	}
}
