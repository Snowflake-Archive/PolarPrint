package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/snowflake-software/polarprint/pkg/presenters"
	"github.com/snowflake-software/polarprint/pkg/utils"
)

func RenderUI() fiber.Handler {
	return func(c *fiber.Ctx) error {
		files, err := utils.GetFilenames()

		if err != nil {
			return c.JSON(presenters.FileListErrorResponse(err))
		}

		return c.Render("index", fiber.Map{
			"files": files,
		})
	}
}
