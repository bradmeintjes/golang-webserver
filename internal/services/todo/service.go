package todo

import (
	"context"
	"go/webserver/internal/database"
	"go/webserver/internal/lib/errors"
	"time"

	"github.com/markbates/pop/nulls"

	"github.com/jmoiron/sqlx"
)

// Service to manage todos
type Service struct {
	db *sqlx.DB
}

// NewService creates a new instance of todo service
func NewService(db *sqlx.DB) *Service {
	return &Service{
		db: db,
	}
}

// CreateTodo will create a new todo
func (s *Service) CreateTodo(ctx context.Context, content string) (*Todo, error) {
	todo := New(content)

	err := database.AsTransaction(s.db, func(tx *sqlx.Tx) error {
		return insert(ctx, tx, todo)
	})
	if err != nil {
		return nil, errors.Internalf("failed to create todo: %s", err)
	}

	return todo, nil
}

// CompleteTodo will mark a todo as complete
func (s *Service) CompleteTodo(ctx context.Context, id int64) (*Todo, error) {
	todo, err := selectByID(ctx, s.db, id)
	if err != nil {
		return nil, errors.Internalf("failed to get todo: %s", err)
	}
	if todo == nil {
		return nil, errors.NotFoundf("todo not found with id %d", id)
	}

	err = database.AsTransaction(s.db, func(tx *sqlx.Tx) error {
		todo.Complete = true
		todo.CompletedOn = nulls.NewTime(time.Now())
		return update(ctx, tx, todo)
	})
	if err != nil {
		return nil, errors.Internalf("failed to update todo: %s", err)
	}
	return todo, nil
}

// DeleteTodo will delete a todo
func (s *Service) DeleteTodo(ctx context.Context, id int64) error {
	err := database.AsTransaction(s.db, func(tx *sqlx.Tx) error {
		return delete(ctx, tx, id)
	})
	if err != nil {
		return errors.Internalf("failed to delete todo: %s", err)
	}
	return nil
}

// FetchAllTodos fetch all the todos
func (s *Service) FetchAllTodos(ctx context.Context) ([]Todo, error) {
	todos, err := selectAll(ctx, s.db)
	if err != nil {
		return nil, errors.Internalf("failed to fetch todos: %s", err)
	}
	return todos, nil
}
