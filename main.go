package main

import (
	"fmt"

	"github.com/abhishek9686/food-order-backend/server"
)

func main() {

	fmt.Println("Starting Food Order Backend Server: 8008")

	// Connect to the Database
	// database.ConnectDB()

	// Listen on PORT 8008
	server.Start()
}
