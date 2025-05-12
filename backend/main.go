package main

import (
	"backend/db"
	"backend/router"
)

func main() {
	db.InitDB("host=localhost port=5432 user=admin password=admin dbname=iptc-editor sslmode=disable")

	server := router.SetupRouter()

	server.Run(":8080")
}
