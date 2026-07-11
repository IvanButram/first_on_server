package crud

import (
	"study/postgres/models"
)

func (crud *CRUD_struct) InsertRow(task models.CreateModel) error {
	sqlQuery := `
		INSERT INTO TASKS (name, description, completed, createdAt)
		VALUES ($1, $2, false, $3);
	`

	_, err := crud.Conn.Exec(crud.Ctx, sqlQuery, task.Title, task.Description, task.CreatedAt)

	return err
}
