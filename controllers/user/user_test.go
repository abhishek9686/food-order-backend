package user_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/abhishek9686/food-order-backend/controllers/user"
	internal_user "github.com/abhishek9686/food-order-backend/internal/user"
	"github.com/abhishek9686/food-order-backend/internal/utils"
)

func TestLogin(t *testing.T) {
	loginReq := internal_user.LoginReq{}
	data, _ := json.Marshal(loginReq)
	req, _ := http.NewRequest(http.MethodPost, "/api/user/login", bytes.NewReader(data))
	w := httptest.NewRecorder()
	user.Login(w, req)
	resp := w.Result()
	defer resp.Body.Close()
	respData, _ := ioutil.ReadAll(resp.Body)
	loginResp := utils.APIResp{}
	json.Unmarshal(respData, &loginResp)
}

func TestSignUp(t *testing.T) {
	signUpreq := internal_user.User{}
	data, _ := json.Marshal(signUpreq)
	req, _ := http.NewRequest(http.MethodPost, "/api/user/signUp", bytes.NewReader(data))
	w := httptest.NewRecorder()
	user.SignUp(w, req)
	resp := w.Result()
	defer resp.Body.Close()
	respData, _ := ioutil.ReadAll(resp.Body)
	signUpresp := utils.APIResp{}
	json.Unmarshal(respData, &signUpresp)

}
