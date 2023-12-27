package common

import (
	"encoding/json"
	"log"
	"net/http"
	"reflect"
)

type response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

type error struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
}

func ResponseJSON(w http.ResponseWriter, code int, payload interface{}) {

	responseType := reflect.TypeOf(payload)

	data, err := json.Marshal(response{
		Success: responseType.Name() != "error",
		Data:    payload,
	})

	if err != nil {
		log.Printf("Marshal failed: %v", err)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func ErrorJSON(w http.ResponseWriter, code int, message string) {
	if code > 499 {
		log.Printf("Server error: %v", message)
	}

	ResponseJSON(w, code, error{
		Status: code,
		Error:  message,
	})
}
