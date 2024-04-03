package ws

import "encoding/json"

type Event struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type EventHandler func(event Event, c *Client) error

const (
	EventUpdatePrinterArray = "sync_printers_array"
)

type UpdatePrinterArrayEvent struct {
	Printers []string `json:"printers"`
}
