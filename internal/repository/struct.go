package repository

import (
	"context"
	"fmt"
	"study/internal/core"

	"github.com/jackc/pgx/v5"
)

type Repository struct {
	Conn   *pgx.Conn
	Ctx    context.Context
	Logger *core.Logger
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
		r.Logger.Error("error on finding id with title", core.Field{Key: "err: ", Value: err})
		return -1, err
	}

	r.Logger.Debug("title to id received", core.Field{Key: "id: ", Value: id})
	return id, nil
}
