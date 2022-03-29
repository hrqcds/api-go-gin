package main

import (
	"github.com/hrqcds/api-go-gin/database"
	"github.com/hrqcds/api-go-gin/routes"
)

func main() {

	database.DB_conn()

	routes.HandleRequests()
}
