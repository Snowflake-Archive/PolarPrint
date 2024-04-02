package handlers

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/snowflake-software/polarprint/pkg/db"
	"github.com/snowflake-software/polarprint/pkg/presenters"
)

func GetClusters() fiber.Handler {
	return func(c *fiber.Ctx) error {
		clusters, err := db.GetClusters()

		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenters.ClusterListFailedResponse(err))
		}

		return c.JSON(presenters.ClusterListSuccessResponse(clusters))
	}
}

func GetCluster() fiber.Handler {
	return func(c *fiber.Ctx) error {
		idParsed, err := strconv.Atoi(c.Params("clusterId"))
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.GenericClusterFailure(err))
		}

		cluster, err := db.GetCluster(int64(idParsed))
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(err)
		}

		return c.JSON(presenters.GenericClusterSuccess(*cluster))
	}
}

func CreateCluster() fiber.Handler {
	return func(c *fiber.Ctx) error {
		cluster, err := db.CreateCluster()

		if err != nil {
			return c.JSON(presenters.GenericClusterFailure(err))
		}

		return c.JSON(presenters.GenericClusterSuccess(*cluster))
	}
}
