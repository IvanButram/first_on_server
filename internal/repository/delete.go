package repository

func (r *Repository) Delete(id int) error {
	sqlQuery := `
	DELETE FROM TASKS
	WHERE id=$1;
	`

	_, err := r.Conn.Exec(r.Ctx, sqlQuery, id)
	return err
}
