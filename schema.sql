DROP TABLE IF EXISTS tasks;

CREATE TABLE tasks (
  id BIGINT PRIMARY KEY,
  title TEXT NOT NULL,
  done BOOLEAN NOT NULL DEFAULT FALSE,
  created_at BIGINT NOT NULL
);

INSERT INTO tasks (id, title, done, created_at) VALUES
(1, 'Test 1', false, 0),
(2, 'Test 2', true, 0);
