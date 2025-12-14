package memdb

import (
  "sync"

  "go-news/pkg/storage"
)

type Store struct {
  mu    sync.RWMutex
  tasks map[int]storage.Task
}

func New() *Store {
  s := &Store{tasks: make(map[int]storage.Task)}
  s.tasks[1] = storage.Task{ID: 1, Title: "Test 1", Done: false, CreatedAt: 0}
  s.tasks[2] = storage.Task{ID: 2, Title: "Test 2", Done: true, CreatedAt: 0}
  return s
}

func (s *Store) Tasks() ([]storage.Task, error) {
  s.mu.RLock()
  defer s.mu.RUnlock()

  out := make([]storage.Task, 0, len(s.tasks))
  for _, t := range s.tasks {
    out = append(out, t)
  }
  return out, nil
}

func (s *Store) AddTask(t storage.Task) error {
  s.mu.Lock()
  defer s.mu.Unlock()

  if _, ok := s.tasks[t.ID]; ok {
    return storage.ErrAlreadyExists
  }
  s.tasks[t.ID] = t
  return nil
}

func (s *Store) UpdateTask(t storage.Task) error {
  s.mu.Lock()
  defer s.mu.Unlock()

  if _, ok := s.tasks[t.ID]; !ok {
    return storage.ErrNotFound
  }
  s.tasks[t.ID] = t
  return nil
}

func (s *Store) DeleteTask(t storage.Task) error {
  s.mu.Lock()
  defer s.mu.Unlock()

  if _, ok := s.tasks[t.ID]; !ok {
    return storage.ErrNotFound
  }
  delete(s.tasks, t.ID)
  return nil
}
