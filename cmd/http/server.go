package http

import (
	"log"
	"net/http"
	"time"

	"github.com/food-order-app/food-order-backend/internal/meals"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/* Todo
1. list meals api
*/

func Start() {

	router := mux.NewRouter()
	router.HandleFunc("/api/meals", meals.GetMeals).Methods("GET")
	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(":8008", handler))

}
func starterMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {

		start := time.Now()
		log.Printf("Duration: %s %s %s", req.Method, req.RequestURI, time.Since(start).String())
	})
}
