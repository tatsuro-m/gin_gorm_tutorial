package handler

import "github.com/gin-gonic/gin"

func BindRoute(r *gin.RouterGroup) {
	r.GET(userIndex, UserIndex())
}
