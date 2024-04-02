package handlers

import (
	"fmt"
	"net/http"
	"os"

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

func GetFile() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendFile("./prints/" + c.Params("fileId"))
	}
}

func UploadFile() fiber.Handler {
	return func(c *fiber.Ctx) error {
		file, err := c.FormFile("file")
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenters.FileUploadFailedResponse(err))
		}

		err = c.SaveFile(file, fmt.Sprintf("./prints/%s", file.Filename))
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenters.FileUploadFailedResponse(err))
		}

		return c.JSON(presenters.FileUploadedResponse("mohahha")) // TODO: Fetch real ID
	}
}

type RenameFileRequestBody struct {
	NewName string `json:"newName"`
}

func RenameFile() fiber.Handler {
	return func(c *fiber.Ctx) error {
		data := new(RenameFileRequestBody)

		if err := c.BodyParser(data); err != nil {
			return c.JSON(presenters.GenericErrorResponse(err))
		}

		os.Rename("./prints/"+c.Params("fileId"), "./prints/"+data.NewName)
		return c.JSON(presenters.GenericOKResponse())
	}
}

func DeleteFile() fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := os.Remove("./prints/" + c.Params("fileId"))

		if err != nil {
			return c.JSON(presenters.GenericErrorResponse(err))
		}

		return c.JSON(presenters.GenericOKResponse())
	}
}
