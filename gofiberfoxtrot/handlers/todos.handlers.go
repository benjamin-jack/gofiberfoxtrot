package handlers

import (
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2"
	"github.com/benjamin-jack/gofiberfoxtrot/views"
	"github.com/benjamin-jack/gofiberfoxtrot/models"
	"github.com/a-h/templ"
	"log"
//	"github.com/jackc/pgx/v5"
//	"context"
)
/*
func getTodos()([]models.Todo){
	var conn = models.Db
	data, _ := conn.Query(context.Background(), "select * from todo order by id asc;")
		list, err := pgx.CollectRows(data, pgx.RowToStructByPos[models.Todo])
	if err != nil { return []models.Todo{}}
	return list
}
*/
func RenderTodos(c *fiber.Ctx) error {
	todos := views.TodoIndex(models.GetTodos())
	handler := adaptor.HTTPHandler(templ.Handler(todos))
	return handler(c)
}

func UpdateTodos(c *fiber.Ctx) error {
	todos := views.TodoList(models.GetTodos())
	handler := adaptor.HTTPHandler(templ.Handler(todos))
	return handler(c)
}
// ADD TODO FUNCTIONS IN MODELS GO FILE
func HandleViewTodosList(c *fiber.Ctx) error {
	//todos := views.TodoIndex(models.GetTodos())
	//handler := adaptor.HTTPHandler(templ.Handler(todos))
	return RenderTodos(c)
}

func HandleViewTodosEdit(c *fiber.Ctx) error {
	if c.Method() == "GET" {
		log.Println("read signal pushed GET")	
	}
	if c.Method() == "POST" {
		models.AddTodo(c.FormValue("create-todo"))
	}
	if c.Method() == "PATCH" {
		models.StatusTodo(c.Query("todo-id"))
	}
	if c.Method() == "DELETE" {
		models.RemoveTodo(c.Query("todo-id"))
	}
	return UpdateTodos(c)
}
