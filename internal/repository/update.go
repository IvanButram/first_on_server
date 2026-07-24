package repository

import "time"

func (r *Repository) Update(id int) error {
	sqlQuery := `
	UPDATE TASKS
	SET completed=true, completedAt=$1
	WHERE id=$2;
	`

	_, err := r.Conn.Exec(r.Ctx, sqlQuery, time.Now(), id)
	return err
}
