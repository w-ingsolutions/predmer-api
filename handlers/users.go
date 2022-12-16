package handlers

import (
	"predmer-api/database"

	"github.com/gofiber/fiber/v2"
)

// UserGet returns a user
func Radovi(c *fiber.Ctx) error {
	radovi := database.Get()

	return c.JSON(fiber.Map{
		"success": true,
		"radovi":  radovi,
	})
}

// // UserCreate registers a user
// func UserCreate(c *fiber.Ctx) error {
// 	user := &models.User{
// 		// Note: when writing to external database,
// 		// we can simply use - Name: c.FormValue("user")
// 		Name: utils.CopyString(c.FormValue("user")),
// 	}
// 	database.Insert(user)

// 	return c.JSON(fiber.Map{
// 		"success": true,
// 		"user":    user,
// 	})
// }

// // NotFound returns custom 404 page
// func NotFound(c *fiber.Ctx) error {
// 	return c.Status(404).SendFile("./static/private/404.html")
// }
