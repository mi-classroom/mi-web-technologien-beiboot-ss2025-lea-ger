package main

import "backend/router"

func main() {
	server := router.SetupRouter()

	server.Run(":8080")
}
