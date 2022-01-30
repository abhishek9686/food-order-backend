package main

import (
	"fmt"

	"github.com/food-order-app/food-order-backend/cmd/http"
)

func main() {

	fmt.Println("Starting Food Order Backend Server: 8008")
	http.Start()
}
