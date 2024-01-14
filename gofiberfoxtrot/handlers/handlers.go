package handlers

import (
//	"github.com/benjamin-jack/gofiberfoxtrot/views"
//	"github.com/benjamin-jack/gofiberfoxtrot/views/partials"
//	"github.com/benjamin-jack/gofiberfoxtrot/models"
	"github.com/gofiber/fiber/v2"
//	"github.com/a-h/templ"
//	"github.com/gofiber/fiber/v2/middleware/adaptor"
//	"github.com/jackc/pgx/v5"

)


func Setup(app *fiber.App) {

	/* Public views */
	app.Get("/", HandleViewHome )

	/* Views that are protected */
	todoApp := app.Group("/todos")
	todoApp.Get("/", HandleViewTodosList)
	todoApp.Get("/edit/", HandleViewTodosEdit)
	todoApp.Post("/edit/", HandleViewTodosEdit)
	todoApp.Delete("/edit/", HandleViewTodosEdit)
	todoApp.Patch("/edit/", HandleViewTodosEdit)
	/* TODO not found manager */
}


