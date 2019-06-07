package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type heartBeatResponse struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}

func heartbeat(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(heartBeatResponse{Status: "OK", Code: 200})
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", heartbeat)
	log.Fatal(http.ListenAndServe(":8080", router))
}
