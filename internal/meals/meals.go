package meals

import (
	"encoding/json"
	"net/http"
)

func GetMeals(w http.ResponseWriter, r *http.Request) {
	resp := mealItems

	out, _ := json.Marshal(resp)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")
	_, _ = w.Write(out)

}
