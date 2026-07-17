package crud

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type CRUD_struct struct {
	Conn *pgx.Conn
	Ctx  context.Context
}

func (crud *CRUD_struct) TitleToID(title string) (int, error) {
	var id int

	sqlQuery := `
	SELECT id FROM TASKS
	WHERE name=$1;
	`
	err := crud.Conn.QueryRow(crud.Ctx, sqlQuery, title).Scan(&id)
	if err != nil {
		fmt.Println("error on finding id with title")
		return -1, err
	}

	return id, nil
}
