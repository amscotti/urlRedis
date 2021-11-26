package handlers

import (
	"html"

	"github.com/amscotti/urlRedis/storage"
	"github.com/gofiber/fiber/v2"
)

func CreateKey(c *fiber.Ctx) error {
	url := c.FormValue("url")

	status, err := storage.DBConn.Set(html.EscapeString(url))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(status)
}

func GetKey(c *fiber.Ctx) error {
	key := c.Params("key")

	status, err := storage.DBConn.Get(key)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString(err.Error())
	}

	return c.JSON(status)
}

func RedirectKey(c *fiber.Ctx) error {
	key := c.Params("key")

	status, err := storage.DBConn.Get(key)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString(err.Error())
	}

	storage.DBConn.Incr(key)
	return c.Redirect(html.EscapeString(status.URL))
}
