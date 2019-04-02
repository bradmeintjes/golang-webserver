package todo

import (
	"context"
	"go/webserver/internal/database"
	"go/webserver/internal/util/testutil"
	"testing"
)

func TestService(t *testing.T) {
	ctx := context.Background()
	db, clean, err := database.Fixture("todo_test.db")
	defer clean()
	testutil.AssertNoErr(t, err, "create db")

	svc := NewService(db)

	testContent := "test todo"

	todo, err := svc.CreateTodo(ctx, testContent)
	testutil.AssertNoErr(t, err, "create todo")
	testutil.AssertNotNil(t, todo, "create todo")
	testutil.CheckEq(t, int64(1), todo.ID, "create todo sets ID")
	testutil.CheckEq(t, false, todo.Complete, "create todo is incomplete")
	testutil.CheckEq(t, testContent, todo.Content, "create todo has content")
	testutil.Check(t, todo.Created.Unix() > 0, "create todo has created date")

	todos, err := svc.FetchAllTodos(ctx)
	testutil.AssertNoErr(t, err, "fetch todos")
	testutil.AssertNotNil(t, todos, "fetch todos")
	testutil.CheckEq(t, 1, len(todos), "fetch todos")
	checkTodoEq(t, *todo, todos[0], "fetch todos")

	todo, err = svc.CompleteTodo(ctx, todo.ID)
	testutil.AssertNoErr(t, err, "complete todo")
	testutil.AssertNotNil(t, todo, "complete todo")
	testutil.CheckEq(t, true, todo.Complete, "complete todo is complete")
	testutil.Check(t, todo.CompletedOn.Valid && todo.CompletedOn.Time.Unix() > 0, "complete todo has date")

	err = svc.DeleteTodo(ctx, todo.ID)
	testutil.AssertNoErr(t, err, "delete todo")
}

func checkTodoEq(t *testing.T, expected Todo, got Todo, task string) {
	testutil.CheckEq(t, expected.ID, got.ID, task)
	testutil.CheckEq(t, expected.Content, got.Content, task)
	testutil.CheckEq(t, expected.Complete, got.Complete, task)
	testutil.CheckEq(t, expected.Created.Unix(), got.Created.Unix(), task)
	if expected.CompletedOn.Valid {
		testutil.CheckEq(t, expected.CompletedOn.Time.Unix(), got.CompletedOn.Time.Unix(), task)
	} else {
		testutil.CheckEq(t, false, got.CompletedOn.Valid, task)
	}
}
