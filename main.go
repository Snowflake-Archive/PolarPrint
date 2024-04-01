package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/handlebars/v2"
	files "github.com/snowflake-software/polarprint/utils"
)

type RenameFileRequestBody struct {
	NewName string `json:"newName"`
}

func main() {
	app := fiber.New(fiber.Config{
		Views:   handlebars.New("./views", ".hbs"),
		Prefork: true,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"files": files.GetFilenames(),
		})
	})

	app.Get("/files", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"files": files.GetFilenames(),
		})
	})

	app.Get("/files/:fileId", func(c *fiber.Ctx) error {
		return c.SendFile("./prints/" + c.Params("fileId"))
	})

	app.Put("/files", func(c *fiber.Ctx) error {
		file, err := c.FormFile("file")
		if err != nil {
			return err
		}

		return c.SaveFile(file, fmt.Sprintf("./prints/%s", file.Filename))
	})

	app.Patch("/files/:fileId", func(c *fiber.Ctx) error {
		data := new(RenameFileRequestBody)

		if err := c.BodyParser(data); err != nil {
			return err
		}

		os.Rename("./prints/"+c.Params("fileId"), data.NewName)
		return c.SendString("OK")
	})

	app.Delete("/files/:fileId", func(c *fiber.Ctx) error {
		os.Remove("./prints/" + c.Params("fileId"))
		return c.SendString("OK")
	})

	app.Get("/queue", func(c *fiber.Ctx) error {
		return c.SendString("Yes")
	})

	app.Post("/queue", func(c *fiber.Ctx) error {
		return c.SendString("Yes")
	})

	app.Delete("/queue/:itemId", func(c *fiber.Ctx) error {
		return c.SendString("Yes")
	})

	app.Static("/", "./static")

	log.Fatal(app.Listen(":3000"))
}
