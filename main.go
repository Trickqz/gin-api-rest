package main

import (
	"github.com/trickqz/api-go-gin/database"
	"github.com/trickqz/api-go-gin/routes"
)

func main() {
	database.ConnectBD()
	routes.HandleRequests()
}
