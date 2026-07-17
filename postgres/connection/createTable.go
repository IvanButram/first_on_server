package connection

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateTableUsers(conn *pgx.Conn, ctx context.Context) error {
	sqlQuery := `
	CREATE TABLE IF NOT EXISTS TASKS (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		description VARCHAR(500) NOT NULL,
		completed BOOLEAN NOT NULL,
		createdAt TIMESTAMP NOT NULL,
		completedAt TIMESTAMP,

		UNIQUE(name)
		 );
	`
	_, err := conn.Exec(ctx, sqlQuery)

	return err
}
