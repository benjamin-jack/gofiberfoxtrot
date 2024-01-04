package main

import (
	"context"
	"fmt"
	"strings"
	"os"
	//"log"
	//"net/http"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/handlebars/v2"
	"github.com/jackc/pgx/v5"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/a-h/templ"
	"github.com/benjamin-jack/gofiberfoxtrot/views"
	"github.com/benjamin-jack/gofiberfoxtrot/models"
)

func getList(conn *pgx.Conn)([]models.Todo) {
	data, _ := conn.Query(context.Background(), "select * from todo order by id asc;")
		sing, err := pgx.CollectRows(data, pgx.RowToStructByPos[models.Todo])
	if err != nil { return []models.Todo{}}
	return sing
	}

func main() {

	databaseURL := "postgres://user123:pass123@db:5432/postgres"
	conn, err := pgx.Connect(context.Background(), databaseURL)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	engine := handlebars.New("./views",".hbs")
	
	app := fiber.New(fiber.Config{
	    Views: engine,
	})

	app.Static("/","./")

	app.Static("/","./scripts")

	app.Get("/", func(c *fiber.Ctx) error {
    		return c.Render("index",fiber.Map{"Todoslist": getList(conn),})
	})

	app.Get("/done", func(c *fiber.Ctx) error {
		tx, err := conn.Begin(context.Background())
		
		defer tx.Rollback(context.Background())

		_, err = tx.Exec(context.Background(), "UPDATE todo SET isdone = NOT isdone WHERE id="+c.Query("todo-id")+";")
		if err != nil { return err }

		err = tx.Commit(context.Background())
		
		return c.Render("todos", fiber.Map{"Todoslist": getList(conn),})
	})

	//templ handler
	app.Get("/test", func(c *fiber.Ctx) error {
		
		todos := views.TodoIndex(getList(conn))
		handler := adaptor.HTTPHandler(templ.Handler(todos))
		return handler(c)

	})
	
	app.Get("/get", func(c *fiber.Ctx) error {
		rows, _ := conn.Query(context.Background(), "select name from todo where isdone=true;")
		names, err := pgx.CollectRows(rows, pgx.RowTo[string])
		if err != nil { return err }
		return c.SendString(strings.Join(names," "))
	})

	app.Get("/set", func(c *fiber.Ctx) error {
		todoname := c.Query("create-todo")
		fmt.Println(todoname)
		if todoname == "" {
			return c.Render("todos", fiber.Map{"Todoslist": getList(conn),})
		}
		
		_, err := conn.CopyFrom(
    			context.Background(),
    			pgx.Identifier{"todo"},
    			[]string{"name","isdone"},
			pgx.CopyFromRows(
				[][]any{
					{c.Query("create-todo"),"false"},
				}),
		)		
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to copy to database: %v", err)
			return err
		}
		return c.Render("todos", fiber.Map{"Todoslist": getList(conn),})
	})

	app.Get("/del", func(c *fiber.Ctx) error {
		tx, err := conn.Begin(context.Background())
		if err != nil {return err}
		defer tx.Rollback(context.Background())
		_, err = tx.Exec(context.Background(), "delete from todo where id="+c.Query("todo-id")+";")
		if err != nil {return err}
		err = tx.Commit(context.Background())
		if err != nil {return err}
		return c.Render("todos", fiber.Map{"Todoslist": getList(conn),})
	})
	app.Listen(":3000")
}
