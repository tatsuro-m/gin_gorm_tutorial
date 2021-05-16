package handler

import "github.com/gin-gonic/gin"

func BindRoute(r *gin.RouterGroup) {
	r.GET(userIndex, UserIndex())
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	v1 := api.Group("/v1")

	BindRoute(v1)
	return r
}
