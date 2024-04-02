package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/snowflake-software/polarprint/pkg/handlers"
)

func ClusterRoutes(app fiber.Router) {
	app.Get("/clusters", handlers.GetClusters())
	app.Get("/clusters/:clusterId", handlers.GetCluster())
	app.Post("/clusters", handlers.CreateCluster())
}
