package handlers

import (
	"time"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var (
	store         *session.Store
	AUTH_KEY      string = "authenticated"
	USER_ID       string = "user_id"
	fromProtected bool   = false
)

func Setup(app *fiber.App) {
	/* Sessions Config */
	store = session.New(session.Config{
		CookieHTTPOnly: true,
		// CookieSecure: true, for https
		Expiration: time.Hour * 1,
	})

	/* Public views */
	app.Get("/", HandleViewHome)
	app.Get("/login", HandleViewLogin)
	app.Post("/login", HandleViewLogin)
	app.Get("/register", HandleViewRegister)
	app.Post("/register", HandleViewRegister)

	/* Protected via AuthMiddleware */
	todoApp := app.Group("/todos", AuthMiddleware)
	todoApp.Get("/", HandleViewTodosList)
	todoApp.Get("/edit/", HandleViewTodosEdit)
	todoApp.Post("/edit/", HandleViewTodosEdit)
	todoApp.Delete("/edit/", HandleViewTodosEdit)
	todoApp.Patch("/edit/", HandleViewTodosEdit)
	/* TODO not found manager */
}


