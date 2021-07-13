package user

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/deepak-v4/bookstore_users-api/domain/users"
	"github.com/deepak-v4/bookstore_users-api/services"
	"github.com/deepak-v4/bookstore_users-api/utils/error"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := error.NewBadRequest("invalid json body")
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

func GetUser(c *gin.Context) {

	userId, userErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if userErr != nil {
		err := error.NewBadRequest("user id should be a number")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, user)

}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me !!")
}
