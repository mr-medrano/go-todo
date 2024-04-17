-- +goose Up
CREATE TABLE tasks (
  id UUID NOT NULL PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  title TEXT NOT NULL,
  note TEXT
);

-- +goose Down
DROP TABLE tasks;