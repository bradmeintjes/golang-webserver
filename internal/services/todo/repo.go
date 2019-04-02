package todo

import (
	"context"
	"go/webserver/internal/util/sqlutil"

	"github.com/jmoiron/sqlx"
)

var (
	insertTodo     = sqlutil.GenerateInsert("todos", Todo{}, []string{"id"})
	updateTodo     = sqlutil.GenerateUpdate("todos", Todo{}, []string{})
	deleteTodo     = "delete from todos where id = ?"
	selectTodos    = "select * from todos"
	selectTodoByID = selectTodos + " where id = ?"
)

// store a todo item
func insert(ctx context.Context, tx *sqlx.Tx, todo *Todo) error {
	res, err := tx.NamedExecContext(ctx, insertTodo, todo)
	if err == nil {
		todo.ID, err = res.LastInsertId()
	}
	return err
}

// update a todo item
func update(ctx context.Context, tx *sqlx.Tx, todo *Todo) error {
	_, err := tx.NamedExecContext(ctx, updateTodo, todo)
	return err
}

// delete removes a todo item
func delete(ctx context.Context, tx *sqlx.Tx, id int64) error {
	_, err := tx.ExecContext(ctx, deleteTodo, id)
	return err
}

// select a todo by id
func selectByID(ctx context.Context, db *sqlx.DB, id int64) (*Todo, error) {
	var todo Todo
	err := db.GetContext(ctx, &todo, selectTodoByID, id)
	return &todo, err
}

// select all todos
func selectAll(ctx context.Context, db *sqlx.DB) ([]Todo, error) {
	todos := make([]Todo, 0)

	rows, err := db.QueryxContext(ctx, selectTodos)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var todo Todo
		err = rows.StructScan(&todo)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}
