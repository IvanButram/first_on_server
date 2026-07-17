package main

import (
	"context"
	"fmt"
	Http "study/http"
	crud "study/postgres/CRUD"
	"study/postgres/connection"
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

	Crud_obj := crud.CRUD_struct{
		Conn: conn,
		Ctx:  ctx,
	}

	//передать Crud в инициализацию хендлеров и передать все хендлеры в сервер
	handlers := Http.NewHandlers(&Crud_obj)
	server := Http.NewServer(handlers)

	server.StartServer()
}
