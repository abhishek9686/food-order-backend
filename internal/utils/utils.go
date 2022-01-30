package utils

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func ReadAndParseInput(w http.ResponseWriter, r *http.Request, input interface{}) error {

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, MaxRestAPIPayload))
	if err != nil {
		log.Printf("Error in Reading request Body %+v", err)
		return err
	}
	if err1 := r.Body.Close(); err1 != nil {
		log.Printf("Error in Closing body %s\n", err1.Error())
		return err1
	}

	if err2 := json.Unmarshal(body, input); err2 != nil {

		err2 = GetUnmarshallErrorString(err2)
		log.Printf("Unmarshalling Error. %+v ", err2)

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		output := APIResp{
			ResponseCode:        ResponseFailed,
			ResponseDescription: err2.Error(),
		}
		if err3 := json.NewEncoder(w).Encode(output); err3 != nil {
			log.Printf("Json Encoding Error. %+v", err3)
		}
		return err2
	}
	return nil
}

func GetUnmarshallErrorString(unMarshalErr error) error {
	if ute, ok := unMarshalErr.(*json.UnmarshalTypeError); ok {
		return errors.New("Input " + ute.Value + " for field " + ute.Field + " is incorrect.")
	}
	return unMarshalErr

}
