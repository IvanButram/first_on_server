package main

import (
	"context"
	"fmt"
	Http "study/internal/http"
	"study/internal/repository"
	connection "study/pkg/postgres"
)

func main() {
	ctx := context.Background()
	conn := connection.CheckConnection(ctx)
	err := connection.CreateTableUsers(conn, ctx)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Table is succesfully created")
	}

	r := repository.Repository{
		Conn: conn,
		Ctx:  ctx,
	}

	//передать Crud в инициализацию хендлеров и передать все хендлеры в сервер
	handlers := Http.NewHandlers(&r)
	server := Http.NewServer(handlers)

	server.StartServer()
}
