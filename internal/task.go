package task

import (
	"github.com/google/uuid"
)

type Task struct {
	ID     string
	Name   string
	Status bool
}

func NewTask(name string) Task {
	return Task{
		ID : uuid.New().String(),
		Name:   name,
		Status: false,
	}
}

