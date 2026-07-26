package main

import (
	"context"
	"fmt"
	"os"
	"study/internal/core"
	Http "study/internal/http"
	"study/internal/repository"
	connection "study/pkg/postgres"
)

func main() {

	err := os.MkdirAll("logs", 0755)
	if err != nil {
		fmt.Println("error on creating direcrory for logger: ", err)
	}
	logger, err := core.NewLogger(true, "logs/app.log")
	if err != nil {
		fmt.Println("logger is not proccesing")
	}

	ctx := context.Background()
	conn := connection.CheckConnection(ctx)
	err = connection.CreateTableUsers(conn, ctx)
	if err != nil {
		logger.Error("error on creating table: ", core.Field{Key: "err: ", Value: err})
		panic(err)
	} else {
		logger.Debug("table is successfully created")
		fmt.Println("Table is succesfully created")
	}

	r := repository.Repository{
		Conn:   conn,
		Ctx:    ctx,
		Logger: logger,
	}

	handlers := Http.NewHandlers(&r, logger)
	server := Http.NewServer(handlers)

	server.StartServer()
}
