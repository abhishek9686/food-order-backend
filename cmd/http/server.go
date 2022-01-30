package http

import (
	"log"
	"net/http"

	"github.com/food-order-app/food-order-backend/internal/meals"
	"github.com/gorilla/mux"
)

/* Todo
1. list meals api
*/

func Start() {

	router := mux.NewRouter()
	router.HandleFunc("/api/meals", meals.GetMeals).Methods("GET")

	log.Fatal(http.ListenAndServe(":8008", router))

}
