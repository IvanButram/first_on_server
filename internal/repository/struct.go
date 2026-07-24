package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Repository struct {
	Conn *pgx.Conn
	Ctx  context.Context
}

func (r *Repository) TitleToID(title string) (int, error) {
	var id int

	sqlQuery := `
	SELECT id FROM TASKS
	WHERE name=$1;
	`
	err := r.Conn.QueryRow(r.Ctx, sqlQuery, title).Scan(&id)
	if err != nil {
		fmt.Println("error on finding id with title")
		return -1, err
	}

	return id, nil
}
