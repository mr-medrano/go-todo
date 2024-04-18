package models

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Task struct {
	ID      string
	Title   string
	Note    string
	Created time.Time
	Updated time.Time
}

type TaskModel struct {
	DB *pgxpool.Pool
}

func (m *TaskModel) Insert(ctx context.Context, title string, note string) (string, error) {
	stmt := `INSERT into tasks (title, note) VALUES (@title, @note) RETURNING id`
	args := pgx.NamedArgs{
		"title": title,
		"note":  note,
	}

	row := m.DB.QueryRow(ctx, stmt, args)

	var t Task

	err := row.Scan(&t.ID)
	if err != nil {
		return "", err
	}

	return t.ID, nil
}
