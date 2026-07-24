package repository

import (
	"study/pkg/postgres/models"
	"time"
)

func (r *Repository) InsertRow(task models.CreateModel) error {
	sqlQuery := `
		INSERT INTO TASKS (name, description, completed, createdAt)
		VALUES ($1, $2, false, $3);
	`

	_, err := r.Conn.Exec(r.Ctx, sqlQuery, task.Title, task.Description, time.Now())

	return err
}
