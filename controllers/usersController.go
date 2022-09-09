package controllers

import (
	"fmt"
	"go_bookstore-api/domain/users"
	"go_bookstore-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err)
		//TODO bad request to caller
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		//TODO handle user creation error
	}
	fmt.Println(user)

	c.JSON(http.StatusCreated, result)
}
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me")
}
func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me")
}
