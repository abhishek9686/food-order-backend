package user

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/abhishek9686/food-order-backend/database"
	"github.com/abhishek9686/food-order-backend/internal/user"
	"github.com/abhishek9686/food-order-backend/internal/utils"
)

// Login - api to login the user.
func Login(w http.ResponseWriter, r *http.Request) {
	var req user.LoginReq
	var err error
	resp := utils.APIResp{
		ResponseCode:        utils.ResponseOk,
		ResponseDescription: "successfully logged in",
	}
	if err := utils.ReadAndParseInput(w, r, &req); err != nil {
		log.Println("Failed to parse Login req: ", err)
		return
	}
	defer func() {
		if err != nil {
			resp.ResponseCode = utils.ResponseFailed
			resp.ResponseDescription = err.Error()
		}
		out, _ := json.Marshal(resp)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		w.Header().Add("Content-Type", "application/json")
		_, _ = w.Write(out)
	}()
	err = user.CheckIfValidUser(database.DB, req)
}

// SignUp - api to add new user to the app.
func SignUp(w http.ResponseWriter, r *http.Request) {
	var req user.User
	var err error
	resp := utils.APIResp{
		ResponseCode:        utils.ResponseOk,
		ResponseDescription: "",
	}
	if err := utils.ReadAndParseInput(w, r, &req); err != nil {
		log.Println("Failed to parse SignUp req: ", err)
		return
	}

	defer func() {
		if err != nil {
			resp.ResponseCode = utils.ResponseFailed
			resp.ResponseDescription = err.Error()
		}
		out, _ := json.Marshal(resp)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		w.Header().Add("Content-Type", "application/json")
		_, _ = w.Write(out)
	}()
	_, err = req.CreateUser(database.DB)

}
