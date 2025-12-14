package storage

import "errors"

var (
  ErrNotFound      = errors.New("not found")
  ErrAlreadyExists = errors.New("already exists")
)

type Task struct {
  ID        int   `json:"id" bson:"id"`
  Title     string `json:"title" bson:"title"`
  Done      bool   `json:"done" bson:"done"`
  CreatedAt int64  `json:"created_at" bson:"created_at"`
}

type Interface interface {
  Tasks() ([]Task, error)
  AddTask(Task) error
  UpdateTask(Task) error
  DeleteTask(Task) error
}
