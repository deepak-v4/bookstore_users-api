package user

import (
	"fmt"
	"net/http"

	"github.com/deepak-v4/bookstore_users-api/domain/users"
	"github.com/deepak-v4/bookstore_users-api/services"
	"github.com/deepak-v4/bookstore_users-api/utils/error"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User
	//fmt.Println("Inside Create user")
	/*data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		//todo
		return
	}
	err = json.Unmarshal(data, &user)
	if err != nil {
		//todo
		return
	}
	*/
	if err := c.ShouldBindJSON(&user); err != nil {
		//todo return bad request to the caller

		restErr := error.RestErr{
			Message: "invalid json body",
			Status:  http.StatusBadRequest,
			Error:   "bad_request",
		}
		c.JSON(restErr.Status, restErr)
		return
	}

	fmt.Println(user)
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		// todo user creation error
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me !!")
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me !!")
}
