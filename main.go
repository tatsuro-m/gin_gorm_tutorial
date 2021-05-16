package main

import (
	"gin_gorm_tutorial/handler"
)

func main() {
	r := handler.SetupRouter()
	r.Run()
}
