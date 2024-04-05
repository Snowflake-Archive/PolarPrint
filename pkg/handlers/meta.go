package handlers

import (
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/snowflake-software/polarprint/pkg/db"
	"github.com/snowflake-software/polarprint/pkg/presenters"
	"github.com/snowflake-software/polarprint/pkg/utils"
)

func RenderUI() fiber.Handler {
	return func(c *fiber.Ctx) error {
		files, err := utils.GetFilenames()

		if err != nil {
			return c.JSON(presenters.FileListErrorResponse(err))
		}

		return c.Render("index", fiber.Map{
			"files": files,
		})
	}
}

func GenerateInstaller() fiber.Handler {
	return func(c *fiber.Ctx) error {
		dat, err := os.ReadFile("./views/installer.lua")
		if err != nil {
			log.Print("Error reading file views/installer.lua!", err.Error())
		}
		content := string(dat)
		host := string(c.BaseURL() + "/")

		return c.SendString(strings.Replace(content, "--{{AUTOPILOT_INJECT}}", "host, autopilot = \""+host+"\", true", 1))
	}
}

func AgentDownload() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authorization := c.Get("Authorization")
		if authorization == "" {
			c.Status(401)
			return c.JSON(presenters.MissingAuthorization())
		}

		cluster, err := db.GetClusterByKey(authorization)
		if err != nil || cluster != nil {
			c.Status(401)
			return c.JSON(presenters.Unauthorized())
		}

		return c.SendFile("./views/agent.lua")
	}
}
