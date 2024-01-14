package main

import (
	"context"
	//"fmt"
	//"strings"
	//"os"
	//"log"
	//"net/http"
	"github.com/gofiber/fiber/v2"
	//"github.com/gofiber/template/handlebars/v2"
	"github.com/jackc/pgx/v5"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/a-h/templ"
	//"github.com/benjamin-jack/gofiberfoxtrot/views"
	"github.com/benjamin-jack/gofiberfoxtrot/models"
	"github.com/benjamin-jack/gofiberfoxtrot/handlers"
)

func getList(conn *pgx.Conn)([]models.Todo) {
	data, _ := conn.Query(context.Background(), "select * from todo order by id asc;")
		sing, err := pgx.CollectRows(data, pgx.RowToStructByPos[models.Todo])
	if err != nil { return []models.Todo{}}
	return sing
	}

func Render(c *fiber.Ctx, component templ.Component, options ...func(*templ.ComponentHandler)) error {
	componentHandler := templ.Handler(component)
	for _, o := range options {
		o(componentHandler)
	}
	return adaptor.HTTPHandler(componentHandler)(c)
func init() {
	models.DatabaseConnect()
}
func main() {
	app := fiber.New(fiber.Config{})
	app.Static("/","./")
	app.Static("/","./scripts")
	handlers.Setup(app)
	app.Listen(":3000")
}
