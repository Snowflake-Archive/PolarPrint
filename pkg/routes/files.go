package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/snowflake-software/polarprint/pkg/handlers"
)

func FilesRouter(app fiber.Router) {
	app.Get("/files", handlers.GetFilenames())
	app.Get("/files/:fileId", handlers.GetFile())
	app.Put("/files", handlers.UploadFile())
	app.Patch("/files/:fileId", handlers.RenameFile())
	app.Delete("/files/:fileId", handlers.DeleteFile())
}
