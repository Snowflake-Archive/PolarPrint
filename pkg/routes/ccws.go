package routes

import (
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/snowflake-software/polarprint/pkg/db"
)

func CCWSRoutes(app fiber.Router) {
	app.Use("/ccws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			return c.Next()
		}

		return fiber.ErrUpgradeRequired
	})

	app.Get("/ccws/:key", websocket.New(func(c *websocket.Conn) {
		if _, err := db.GetClusterByKey(c.Params("key")); err != nil {
			log.Print("Error while verifying ccws key:", err)
			c.Close()
			return
		}

		var (
			mt  int
			msg []byte
			err error
		)

		for {
			if mt, msg, err = c.ReadMessage(); err != nil {
				log.Println("error while reading:", err)
				c.Close()
				return
			}

			log.Printf("Received: %s", msg)

			if err = c.WriteMessage(mt, msg); err != nil {
				log.Print("error while writing:", err)
				c.Close()
				return
			}
		}
	}))
}
