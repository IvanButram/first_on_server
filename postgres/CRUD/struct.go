package crud

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type CRUD_struct struct {
	Conn *pgx.Conn
	Ctx  context.Context
}
