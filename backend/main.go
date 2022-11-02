package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type RequestJson struct {
	RequestMessage string `json:"message"`
}

type ResponseJson struct {
	ResponseMessage string `json:"message"`
}

func hoge(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	var requestJson RequestJson
	err := json.NewDecoder(r.Body).Decode(&requestJson)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var responseMessage string
	if requestJson.RequestMessage == "hello" {
		responseMessage = "hello!"
	} else if requestJson.RequestMessage == "bye" {
		responseMessage = "good bye!"
	} else {
		responseMessage = "what?"
	}
	responseJson := ResponseJson{}
	responseJson.ResponseMessage = responseMessage
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseJson)
}

func main() {
	http.HandleFunc("/", hoge)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
