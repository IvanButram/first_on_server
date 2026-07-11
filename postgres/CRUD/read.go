package crud

import "study/postgres/models"

func (crud *CRUD_struct) Read() ([]models.ReadModel, error) {
	sqlQuery := `
	SELECT id, name, description, completed, createdAt, completedAt FROM TASKS;
	`

	var tasks []models.ReadModel

	rows, err := crud.Conn.Query(crud.Ctx, sqlQuery)
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
