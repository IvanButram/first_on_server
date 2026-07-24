package connection

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func CheckConnection(ctx context.Context) *pgx.Conn {
	conn, err := pgx.Connect(ctx, os.Getenv("CONN_STR"))
	if err != nil {
		panic(err)
	}

	err = conn.Ping(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("all is connected")
	return conn
}
