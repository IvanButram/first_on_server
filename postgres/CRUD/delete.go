package crud

func (crud *CRUD_struct) Delete(id int) error {
	sqlQuery := `
	DELETE FROM TASKS
	WHERE id=$1;
	`

	_, err := crud.Conn.Exec(crud.Ctx, sqlQuery, id)
	return err
}
