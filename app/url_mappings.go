package app

import (
	"github.com/deepak-v4/bookstore_users-api/controllers/ping"
	"github.com/deepak-v4/bookstore_users-api/controllers/user"
)

func mapUrls() {

	router.GET("/ping", ping.Ping)
	router.POST("/users", user.Create)
	router.GET("/users/:id", user.Get)
	router.PUT("/users/:id", user.Update)
	router.PATCH("/users/:id", user.Update)
	router.DELETE("/users/:id", user.Delete)
	router.GET("/internal/users/search", user.Search)
	router.POST("/users/login", user.Login)

}
