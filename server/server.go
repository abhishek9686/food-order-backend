package server

import (
	"log"
	"net/http"
	"time"

	controllers_user "github.com/abhishek9686/food-order-backend/controllers/user"
	"github.com/abhishek9686/food-order-backend/internal/meals"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Start() {

	router := mux.NewRouter()
	router.HandleFunc("/api/meals", meals.GetMeals).Methods(http.MethodGet)
	router.HandleFunc("/api/user/login", controllers_user.Login).Methods(http.MethodPost)
	router.HandleFunc("/api/user/signUp", controllers_user.SignUp).Methods(http.MethodPost)
	router.HandleFunc("/api/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte("ok"))
	})
	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(":8008", handler))

}

func starterMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		start := time.Now()
		log.Printf("Duration: %s %s %s", req.Method, req.RequestURI, time.Since(start).String())
	})
}
