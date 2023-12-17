package main

import (
	"log"
	//"net/http"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/handlebars/v2"
	"github.com/gofiber/storage/postgres/v3"	
)
func main() {
	
	engine := handlebars.New("./views",".hbs")

    app := fiber.New(fiber.Config{
	    Views: engine,
    })

    app.Static("/","./")

    app.Static("/","./scripts")

    database := postgres.New(postgres.Config{
	Database: "fiber",
	Username: "benjamin",
	Password: "honeyrose",
    })
//	database := postgres.New()
    var s string = "Ben"

	sb := []byte(s)

	database.Set("name", sb, 0)

	var x, _ = database.Get("name")
	var y string = string(x)

    app.Get("/", func(c *fiber.Ctx) error {
		// return c.SendString("Hello, World 👋!")
    		return c.Render("index",
			fiber.Map{"Title": "Hello, World!",})
	})
	
	app.Get("/get", func(c *fiber.Ctx) error {
  		// return c.Render("results", fiber.Map{"Results", "TEST"	})
		return c.SendString(y)
	})

    log.Fatal(app.Listen(":3000"))
}
