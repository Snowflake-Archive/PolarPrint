package presenters

import (
	"github.com/gofiber/fiber/v2"
	"github.com/snowflake-software/polarprint/pkg/types"
)

func ClusterListSuccessResponse(data []types.Cluster) *fiber.Map {
	return &fiber.Map{
		"clusters": data,
		"error":    nil,
	}
}

func ClusterListFailedResponse(err error) *fiber.Map {
	return &fiber.Map{
		"clusters": nil,
		"error":    err.Error(),
	}
}

func GenericClusterSuccess(data types.Cluster) *fiber.Map {
	return &fiber.Map{
		"cluster": data,
		"error":   nil,
	}
}

func GenericClusterFailure(err error) *fiber.Map {
	return &fiber.Map{
		"cluster": nil,
		"error":   err.Error(),
	}
}
