package handlers

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/snowflake-software/polarprint/pkg/db"
	"github.com/snowflake-software/polarprint/pkg/presenters"
	"github.com/snowflake-software/polarprint/pkg/utils"
)

func GetQueue() fiber.Handler {
	return func(c *fiber.Ctx) error {
		queue, err := db.GetQueue()

		if err != nil {
			c.Status(500)
			return c.JSON(presenters.GetQueueFailedResponse(err))
		}

		return c.JSON(presenters.GetQueueSuccessResponse(queue))
	}
}

type QueuePrintRequestBody struct {
	FilePath string `json:"file"`
	Quantity int    `json:"quantity"`
}

func AddToQueue() fiber.Handler {
	return func(c *fiber.Ctx) error {
		data := new(QueuePrintRequestBody)

		if err := c.BodyParser(data); err != nil {
			return err
		}

		id := time.Now().Unix() + utils.RandomRange(11111111, 99999999)

		_, err := db.DB.Exec(`INSERT INTO queue(id, file, quantity) values(?, ?, ?)`,
			id,
			"./prints/"+data.FilePath,
			data.Quantity,
		)

		if err != nil {
			log.Fatal(err)
		}

		return c.JSON(fiber.Map{
			"id": id,
		})
	}
}

func DeleteOrder() fiber.Handler {
	return func(c *fiber.Ctx) error {
		_, err := db.DB.Exec(`DELETE`) // TODO: fix

		if err != nil {
			log.Print("Error while deleting order", err.Error())
			return c.SendStatus(404)
		}

		return c.SendString("Yes")
	}
}
