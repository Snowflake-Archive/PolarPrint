package ws

import (
	"encoding/json"
	"errors"
	"log"
	"sync"

	"github.com/snowflake-software/polarprint/pkg/handlers"
)

var (
	ErrEventNotSupported = errors.New("this event type is not supported")
)

type Manager struct {
	clients ClientList
	sync.RWMutex
	handlers map[string]EventHandler
}

func NewManager() *Manager {
	m := &Manager{
		clients:  make(ClientList),
		handlers: make(map[string]EventHandler),
	}

	m.setupEventHandler()

	return m
}

func (m *Manager) setupEventHandler() {
	m.handlers[EventUpdatePrinterArray] = func(e Event, c *Client) error {
		var parsed UpdatePrinterArrayEvent
		err := json.Unmarshal(e.Payload, &parsed)
		if err != nil {
			log.Println("error unmarshalling message", err)
		}

		handlers.SyncPrintersArray(parsed.Printers)
		return nil
	}
}

func (m *Manager) routeEvent(event Event, c *Client) error {
	if handler, ok := m.handlers[event.Type]; ok {
		if err := handler(event, c); err != nil {
			return err
		}

		return nil
	}

	return ErrEventNotSupported
}

func (m *Manager) AddClient(client *Client) {
	m.Lock()
	defer m.Unlock()

	m.clients[client] = true
}

func (m *Manager) RemoveClient(client *Client) {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.clients[client]; ok {
		client.connection.Close()
		delete(m.clients, client)
	}
}
