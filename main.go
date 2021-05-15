package main

import (
	"gin_gorm_tutorial/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	api := r.Group("/api")
	v1 := api.Group("/v1")
	handler.BindRoute(v1)

	r.Run()
}
