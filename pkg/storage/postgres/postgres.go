package postgres

import (
  "database/sql"

  "go-news/pkg/storage"

  _ "github.com/lib/pq"
)

type DB struct {
  db *sql.DB
}

func New(connstr string) (*DB, error) {
  db, err := sql.Open("postgres", connstr)
  if err != nil {
    return nil, err
  }
  if err := db.Ping(); err != nil {
    return nil, err
  }
  return &DB{db: db}, nil
}

func (p *DB) Tasks() ([]storage.Task, error) {
  rows, err := p.db.Query(`SELECT id, title, done, created_at FROM tasks ORDER BY id`)
  if err != nil {
    return nil, err
  }
  defer rows.Close()

  var tasks []storage.Task
  for rows.Next() {
    var t storage.Task
    if err := rows.Scan(&t.ID, &t.Title, &t.Done, &t.CreatedAt); err != nil {
      return nil, err
    }
    tasks = append(tasks, t)
  }
  return tasks, rows.Err()
}

func (p *DB) AddTask(t storage.Task) error {
	_, err := p.db.Exec(
		`INSERT INTO tasks (title, done, created_at)
		 VALUES ($1,$2,$3)`,
		t.Title, t.Done, t.CreatedAt,
	)
	return err
}

func (p *DB) UpdateTask(t storage.Task) error {
  res, err := p.db.Exec(
    `UPDATE tasks SET title=$1, done=$2, created_at=$3 WHERE id=$4`,
    t.Title, t.Done, t.CreatedAt, t.ID,
  )
  if err != nil {
    return err
  }
  aff, _ := res.RowsAffected()
  if aff == 0 {
    return storage.ErrNotFound
  }
  return nil
}

func (p *DB) DeleteTask(t storage.Task) error {
  res, err := p.db.Exec(`DELETE FROM tasks WHERE id=$1`, t.ID)
  if err != nil {
    return err
  }
  aff, _ := res.RowsAffected()
  if aff == 0 {
    return storage.ErrNotFound
  }
  return nil
}
