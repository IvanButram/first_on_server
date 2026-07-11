package crud

import "time"

func (crud *CRUD_struct) Update(id int) error {
	sqlQuery := `
	UPDATE TASKS
	SET completed=true, completedAt=$1
	WHERE id=$2;
	`

	_, err := crud.Conn.Exec(crud.Ctx, sqlQuery, time.Now(), id)
	return err
}
