package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/jet/v2"
	"log"
)

func main() {
	viewEngine := jet.New("./views", ".jet.html")
	viewEngine.Reload(true)

	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Dusan Antonijevic",
		AppName:       "Dusan Antonijevic",
		Views:         viewEngine,
		ViewsLayout:   "layouts/app",
	})

	app.Use(logger.New())
	app.Use(recover2.New())

	app.Static("/assets", "./static")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("pages/home", fiber.Map{})
	})

	log.Fatal(app.Listen(":3000"))
}
