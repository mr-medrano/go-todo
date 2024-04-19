package schemas

import "time"

type Task struct {
	ID      string    `json:"id" db:"id"`
	Title   string    `json:"title" db:"id" binding:"required"`
	Note    string    `json:"note" db:"note"`
	Created time.Time `json:"created_at"`
	Updated time.Time `json:"updated_at"`
}
