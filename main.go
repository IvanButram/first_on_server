package main

import (
	"context"
	"fmt"
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

	Crud := crud.CRUD_struct{
		Conn: conn,
		Ctx:  ctx,
	}

	/*task1 := models.CreateModel{
		Title:       "Обед",
		Description: "Победат",
		CreatedAt:   time.Now(),
	}

	err = Crud.InsertRow(task1)
	if err != nil {
		panic(err)
	}*/

	/*err = Crud.Update(2)
	if err != nil {
		panic(err)
	}*/

	tasks, err := Crud.Read()
	if err != nil {
		panic(err)
	}

	fmt.Println(tasks)
}
