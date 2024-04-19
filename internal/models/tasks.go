package models

import (
	"context"
	"fmt"
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

func (m *TaskModel) Get(ctx context.Context, id string) (*Task, error) {
	stmt := `SELECT id, title, note, created_at, updated_at FROM tasks WHERE id::text = @id`
	args := pgx.NamedArgs{
		"id": id,
	}
	fmt.Println(args)
	row := m.DB.QueryRow(ctx, stmt, args)

	var t Task

	err := row.Scan(&t.ID, &t.Title, &t.Note, &t.Created, &t.Updated)
	if err != nil {
		return nil, err
	}

	return &t, nil
}
