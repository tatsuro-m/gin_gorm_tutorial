package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

const users = "/users"
const userIndex = users

func UserIndex() gin.HandlerFunc {
	users := []User{{
		Name:        "テスト1",
		Email:       "test1@example.com",
		PhoneNumber: "11-1111-1111",
	}}

	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": users,
		})
	}
}
