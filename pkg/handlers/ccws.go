package handlers

import "github.com/gofiber/fiber/v2"

func SyncPrintersArray(printers []string) *fiber.Map {
	return &fiber.Map{
		"event": "sync_printers_finish",
		"error": nil,
	}
}
