package app

import (
	"github.com/parsaakbari1209/go-mongo-CRUD-web/controllers/ping"
	"github.com/parsaakbari1209/go-mongo-CRUD-web/controllers/users"
)

func MapUrls() {
	router.GET("/ping", ping.Ping)
	router.GET("/users/find", users.FindUser)
	router.POST("/users/create", users.CreateUser)
}
