package routes

import (
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/snowflake-software/polarprint/pkg/db"
	"github.com/snowflake-software/polarprint/pkg/ws"
)

func CCWSRoutes(app fiber.Router) {
	app.Use("/ccws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			return c.Next()
		}

		return fiber.ErrUpgradeRequired
	})

	manager := ws.NewManager()

	app.Get("/ccws/:key", websocket.New(func(c *websocket.Conn) {
		if _, err := db.GetClusterByKey(c.Params("key")); err != nil {
			log.Print("Error while verifying ccws key:", err)
			c.Close()
			return
		}

		client := ws.NewClient(c, manager)
		manager.AddClient(client)

		client.ReadMessages()
	}))
}
