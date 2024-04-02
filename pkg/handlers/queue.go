package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/snowflake-software/polarprint/pkg/db"
	"github.com/snowflake-software/polarprint/pkg/presenters"
	"github.com/snowflake-software/polarprint/pkg/types"
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

func AddToQueue() fiber.Handler {
	return func(c *fiber.Ctx) error {
		data := new(types.QueuePrintRequestBody)

		if err := c.BodyParser(data); err != nil {
			return err
		}

		id := time.Now().Unix() + utils.RandomRange(11111111, 99999999)
		db.InsertOrder(id, *data)

		return c.JSON(fiber.Map{
			"id": id,
		})
	}
}

func DeleteOrder() fiber.Handler {
	return func(c *fiber.Ctx) error {
		idParsed, err := strconv.Atoi(c.Params("orderId"))

		if err != nil {
			c.Status(http.StatusBadRequest)
			c.JSON(presenters.GenericErrorResponse(err))
		}

		err = db.DeleteOrder(int64(idParsed))

		if err != nil {
			if err.Error() == "order not found" {
				return c.SendStatus(http.StatusNotFound)
			} else {
				c.Status(http.StatusInternalServerError)
				return c.JSON(presenters.GenericErrorResponse(err))
			}
		}

		return c.JSON(presenters.GenericOKResponse())
	}
}
