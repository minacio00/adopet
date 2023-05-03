package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/minacio00/adopet/helpers"
)

// check if the user is authenticated before delete
func Authenticated(c *fiber.Ctx) error {
	method := c.Method()
	if method != "DELETE" {
		return c.Next()
	}
	token := c.GetReqHeaders()["Authorization"]
	if len(token) == 0 {
		return c.Status(fiber.StatusUnauthorized).SendString("missing token")
	}
	token = strings.Replace(token, "Bearer ", "", 1)
	role, err := helpers.ValidateAuth(token)
	if err != nil {
		return err
	}
	if role == "abrigo" {
		return c.Next()
	}
	return c.Status(fiber.StatusUnauthorized).SendString("missing token")
}
