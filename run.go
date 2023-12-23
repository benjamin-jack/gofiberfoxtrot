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
)

func getList(conn *pgx.Conn)(map[int]string) {
	ret := map[int]string{}
	//var todos []string
	//var indices []int
	rows, _ := conn.Query(context.Background(), "select name from todo;")
		names, err := pgx.CollectRows(rows, pgx.RowTo[string])
	indexrows, _ := conn.Query(context.Background(), "select id from todo;")
		ids, err := pgx.CollectRows(indexrows, pgx.RowTo[int])
	if err != nil { return map[int]string{}}
	for i:= 0; i<len(names); i++ {
		//todos = append(todos, names[i])
		//indices = append(indices, ids[i])
		ret[ids[i]] = names[i]
		}
	fmt.Println(ret)
	return ret
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
	app.Static("/","/views/partials/style.css")

	app.Static("/","./scripts")

	app.Get("/", func(c *fiber.Ctx) error {
    		return c.Render("index",fiber.Map{"Todoslist": getList(conn),})	
	})
	
	app.Get("/get", func(c *fiber.Ctx) error {
		rows, _ := conn.Query(context.Background(), "select name from todo where isdone=true;")
		names, err := pgx.CollectRows(rows, pgx.RowTo[string])
		if err != nil { return err }
		return c.SendString(strings.Join(names," "))
	})

	app.Get("/set", func(c *fiber.Ctx) error {
		todoname := c.Query("create-todo")
		if todoname == "" {
			return c.Render("todos", fiber.Map{"Todoslist": getList(conn),})
		}
		rows := [][]any{
			{todoname,"true"},
		}
		
		_, err := conn.CopyFrom(
    			context.Background(),
    			pgx.Identifier{"todo"},
    			[]string{"name","isdone"},
    			pgx.CopyFromRows(rows),
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
