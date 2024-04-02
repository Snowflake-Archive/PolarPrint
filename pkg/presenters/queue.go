package presenters

import (
	"github.com/gofiber/fiber/v2"
	"github.com/snowflake-software/polarprint/pkg/types"
)

func GetQueueSuccessResponse(queue []types.QueueItem) *fiber.Map {
	return &fiber.Map{
		"err":   nil,
		"queue": queue,
	}
}

func GetQueueFailedResponse(err error) *fiber.Map {
	return &fiber.Map{
		"err":   err,
		"queue": nil,
	}
}
