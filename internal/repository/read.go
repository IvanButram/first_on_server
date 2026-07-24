package repository

import "study/pkg/postgres/models"

func (r *Repository) Read() ([]models.ReadModel, error) {
	sqlQuery := `
	SELECT id, name, description, completed, createdAt, completedAt FROM TASKS;
	`

	var tasks []models.ReadModel

	rows, err := r.Conn.Query(r.Ctx, sqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var t models.ReadModel
		err = rows.Scan(&t.Id, &t.Title, &t.Description, &t.Completed, &t.CreatedAt, &t.CompletedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	return tasks, nil
}

func (r *Repository) ReadOne(id int) models.ReadModel {
	sqlQuery := `
	SELECT id, name, description, completed, createdAt, completedAt FROM TASKS
	WHERE id=$1;
	`

	var task models.ReadModel
	r.Conn.QueryRow(r.Ctx, sqlQuery, id).Scan(&task.Id, &task.Title, &task.Description, &task.Completed, &task.CreatedAt, &task.CompletedAt)

	return task
}
