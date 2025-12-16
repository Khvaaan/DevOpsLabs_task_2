DROP TABLE IF EXISTS tasks;

CREATE TABLE tasks (
  id BIGSERIAL PRIMARY KEY,
  title TEXT NOT NULL,
  done BOOLEAN NOT NULL DEFAULT FALSE,
  created_at BIGINT NOT NULL
);

INSERT INTO tasks (title, done, created_at) VALUES
('Test 1', false, 0),
('Test 2', true, 0);

