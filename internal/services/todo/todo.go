package todo

import (
	"errors"
	"time"

	"github.com/markbates/pop/nulls"
)

// Todo represents a single todo item
type Todo struct {
	ID          int64      `json:"id" db:"id"`
	Content     string     `json:"content" db:"content"`
	Complete    bool       `json:"complete" db:"complete"`
	Created     time.Time  `json:"created" db:"created"`
	CompletedOn nulls.Time `json:"completedOn" db:"completed_on"`
}

// New creates a new todo item
func New(content string) *Todo {
	return &Todo{
		Content:  content,
		Complete: false,
		Created:  time.Now(),
	}
}

// Validate will validate an instance of todo
func Validate(todo *Todo) error {
	if todo == nil {
		return errors.New("todo is nil")
	}

	if todo.Content == "" {
		return errors.New("todo has no content")
	}

	return nil
}
