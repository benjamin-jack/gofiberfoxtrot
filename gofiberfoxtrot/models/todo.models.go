package models

import (
	"github.com/jackc/pgx/v5"
	"context"
	"fmt"
)

type Todo struct {
	Id 	int
	Name 	string
	Status 	bool
}

func dbEditTodo(statement string)(error){
	var err error
	tx, err := db.Begin(context.Background())
		if err != nil { return err }
	defer tx.Rollback(context.Background())
	_, err = tx.Exec(context.Background(), statement)
		if err != nil { return err }
	err = tx.Commit(context.Background())
		if err != nil { return err }	
	return nil
}

// LEGACY CODE
/*
func getTodos()([]models.Todo){
	var conn = models.Db
	data, _ := conn.Query(context.Background(), "select * from todo order by id asc;")
		list, err := pgx.CollectRows(data, pgx.RowToStructByPos[models.Todo])
	if err != nil { return []models.Todo{}}
	return list
}
*/

func GetTodos()([]Todo){
	data, _ := db.Query(context.Background(), "select * from todo order by id asc;")
		list, err := pgx.CollectRows(data, pgx.RowToStructByPos[Todo])
	if err != nil { return []Todo{} }
	return list
}

func RemoveTodo(id string){
	if id == "" { return }
	dbEditTodo(fmt.Sprintf("delete from todo where id="+id+";"))
}

func StatusTodo(id string){
	if id == "" { return }
	dbEditTodo(fmt.Sprintf("UPDATE todo SET isdone = NOT isdone WHERE id ="+id+";"))
}

func AddTodo(name string){
	if name == "" { return }
	dbEditTodo(fmt.Sprintf("INSERT INTO todo (name, isdone) VALUES ('%s', false)", name))
}

