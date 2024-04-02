package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/snowflake-software/polarprint/pkg/handlers"
)

func QueueRoutes(app fiber.Router) {
	app.Get("/queue", handlers.GetQueue())
	app.Put("/queue", handlers.AddToQueue())
	app.Delete("/queue/:orderId", handlers.DeleteOrder())
}
