package main

import (
	//"context"
	//"fmt"
	//"strings"
	//"os"
	//"log"
	//"net/http"
	"github.com/gofiber/fiber/v2"
	//"github.com/gofiber/template/handlebars/v2"
	//"github.com/jackc/pgx/v5"
	//"github.com/gofiber/fiber/v2/middleware/adaptor"
	//"github.com/a-h/templ"
	//"github.com/benjamin-jack/gofiberfoxtrot/views"
	"github.com/benjamin-jack/gofiberfoxtrot/models"
	"github.com/benjamin-jack/gofiberfoxtrot/handlers"
)

func init() {
	models.DatabaseMigrate()
}

func main() {
	app := fiber.New(fiber.Config{})
	app.Static("/","./")
	app.Static("/","./scripts")
	handlers.Setup(app)
	app.Listen(":3000")
}
