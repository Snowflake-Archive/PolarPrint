package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/snowflake-software/polarprint/pkg/presenters"
	"github.com/snowflake-software/polarprint/pkg/utils"
)

func GetFilenames() fiber.Handler {
	return func(c *fiber.Ctx) error {
		filenames, err := utils.GetFilenames()

		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenters.FileListErrorResponse(err))
		}

		return c.JSON(presenters.FileListResponse(filenames))
	}
}
