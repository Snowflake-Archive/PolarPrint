package main

import (
	"fmt"
	"log"
	"os"

	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/handlebars/v2"
	_ "github.com/mattn/go-sqlite3"
	"github.com/snowflake-software/polarprint/utils"
)

type RenameFileRequestBody struct {
	NewName string `json:"newName"`
}

func main() {
	utils.PrintWelcome()
	log.Print("Setting up database...")
	db, err := sql.Open("sqlite3", "./polarprint.db")

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS queue(
			id int not null primary key,
			type varchar(255),
			amount int 
		);
	`)

	if err != nil {
		log.Fatal(err)
	}

	log.Print("Successfully initiated database!")

	count := 0
	err = db.QueryRow("SELECT COUNT(*) FROM queue").Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("There are %d queued items.", count)

	app := fiber.New(fiber.Config{
		Views:   handlebars.New("./views", ".hbs"),
		Prefork: true,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"files": utils.GetFilenames(),
		})
	})

	app.Get("/files", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"files": utils.GetFilenames(),
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
