package todo

import "time"

// Todo represents a single todo item
type Todo struct {
	ID          string
	Content     string
	Complete    bool
	Created     time.Time
	CompletedOn *time.Time
}

// New creates a new todo item
func New(content string) *Todo {
	return &Todo{
		ID:          "",
		Content:     content,
		Complete:    false,
		Created:     time.Now(),
		CompletedOn: nil,
	}
}
