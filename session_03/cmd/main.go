package main

import (
	"go-programming-secure-your-go-apps/session_03/database"
	"go-programming-secure-your-go-apps/session_03/router"
)

func main() {
	database.StartDB()

	PORT := ":4000"
	router.StartServer().Run(PORT)
}
