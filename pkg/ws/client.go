package ws

import (
	"encoding/json"
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/snowflake-software/polarprint/pkg/presenters"
)

type Client struct {
	connection *websocket.Conn
	manager    *Manager
}

type ClientList map[*Client]bool

func NewClient(con *websocket.Conn, manager *Manager) *Client {
	return &Client{
		connection: con,
		manager:    manager,
	}
}

func (c *Client) ReadMessages() {
	defer func() {
		log.Println("removing client from manager")
		c.manager.RemoveClient(c)
	}()

	for {
		messageType, payload, err := c.connection.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error reading message: %v", err)
			}

			println("cleaning up ws due to: ", err.Error(), err, messageType, payload)
			break
		}

		log.Println("Payload: ", payload)

		var request Event
		if err := json.Unmarshal(payload, &request); err != nil {
			log.Printf("error unmarshalling message: %v", err)
			break
		}

		if err := c.manager.routeEvent(request, c); err != nil {
			log.Println("error handeling message: ", err)

			if err == ErrEventNotSupported {
				c.SendMessage(presenters.CcwsUnrecognizedPacket(request.Type))
			} else {
				c.SendMessage(presenters.CcwsUnknownError(err))
			}
		}
	}
}

func (c *Client) SendMessage(message *fiber.Map) {
	content, err := json.Marshal(message)
	if err != nil {
		log.Println("error marshalling server --> client packet: ", err)
	}

	if err := c.connection.WriteMessage(websocket.TextMessage, content); err != nil {
		log.Println("error during ws write: ", err)
	}
}
