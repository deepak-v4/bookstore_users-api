package user

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/deepak-v4/bookstore_users-api/domain/users"
	"github.com/deepak-v4/bookstore_users-api/services"
	"github.com/deepak-v4/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {

	userId, userErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequest("user id should be a number")
		c.JSON(err.Status, err)
		return
	}

	fmt.Println("Reached here 4")
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		resErr := errors.NewBadRequest("invalid json body")
		c.JSON(resErr.Status, resErr)
		return
	}

	user.Id = userId
	fmt.Println(user)

	isPartial := c.Request.Method == http.MethodPatch

	result, updateErr := services.UpdateUser(isPartial, user)
	if updateErr != nil {
		c.JSON(updateErr.Status, updateErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func Create(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequest("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	fmt.Println(user)
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func Get(c *gin.Context) {

	userId, userErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequest("user id should be a number")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("x-Public") == "true"))

}

func Delete(c *gin.Context) {

	userId, userErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequest("user id should be a number")
		c.JSON(err.Status, err)
		return
	}

	getErr := services.DeleteUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "Deleted"})

}

func Search(c *gin.Context) {

	usr_status := c.Query("status")

	user, err := services.SearchByStatus(usr_status)
	if err != nil {
		err := errors.NewBadRequest("user id should be a number")
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func Login(c *gin.Context) {

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequest("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.LoginUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusOK, result)

}
