package controllers

import (
	"fmt"
	"go_bookstore-api/domain/users"
	"go_bookstore-api/services"
	"go_bookstore-api/utils/errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserId(userIdParam string) (int64, *errors.RestErr) {
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, errors.NewBadRequestError("user id should be a number")
	}
	return userId, nil
}

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		//TODO bad request to caller
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		//TODO handle user creation error
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	// userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	// if userErr != nil {
	// 	err := errors.NewBadRequestError("user id should be a number")
	// 	c.JSON(err.Status, err)
	// 	return
	// }
	userId, idErr := GetUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		fmt.Println(idErr)
		return
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	//case1
	// userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	// if userErr != nil {
	// 	err := errors.NewBadRequestError("user id should be a number")
	// 	c.JSON(err.Status, err)
	// 	return
	// }
	//case2
	userId, idErr := GetUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		fmt.Println(idErr)
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	user.Id = userId

	//check for update email only
	isPartial := c.Request.Method == http.MethodPatch

	result, err := services.UpdateUser(isPartial, user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func DeleteUser(c *gin.Context) {
	userId, idErr := GetUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		fmt.Println(idErr)
		return
	}

	if err := services.DeleteUser(userId); err != nil {
		c.JSON(err.Status, err)
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
	// c.String(http.StatusOK, "deleted")
}
