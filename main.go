package main

import (
	"gin_gorm_tutorial/db"
	"gin_gorm_tutorial/handler"
)

func main() {
	db.ConnectDB()

	r := handler.SetupRouter()
	err := r.Run()
	if err != nil {
		panic(err)
	}

	db.Close()
}
