package middlewares

import (
	"github.com/ExeCiety/be-presensi-comindo/utils"

	"github.com/gofiber/fiber/v2"
)

func AcceptLanguage(c *fiber.Ctx) error {
	lang := c.Get("Accept-Language")

	if lang == "" {
		lang = "id"
	}

	utils.SetLocale(lang)
	return c.Next()
}
