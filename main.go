package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/handlebars/v2"
	_ "github.com/mattn/go-sqlite3"
	"github.com/snowflake-software/polarprint/pkg/db"
	"github.com/snowflake-software/polarprint/pkg/routes"
	"github.com/snowflake-software/polarprint/pkg/utils"
)

type QueueItem struct {
	id       int
	file     string
	quantity int
}

func main() {
	utils.PrintWelcome()
	sql, err := db.SetupDatabase()
	db.DB = sql

	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New(fiber.Config{
		Views: handlebars.New("./views", ".hbs"),
		// Prefork: true,
	})

	routes.MetaRouter(app)
	routes.FilesRouter(app)
	routes.QueueRoutes(app)

	app.Static("/", "./static")

	log.Fatal(app.Listen(":3000"))
}
