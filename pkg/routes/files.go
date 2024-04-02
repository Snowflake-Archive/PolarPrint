package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/snowflake-software/polarprint/pkg/handlers"
)

func FilesRouter(app fiber.Router) {
	app.Get("/files", handlers.GetFilenames())
}
