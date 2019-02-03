package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type OrderBy struct {
	Field string `json:"field"`
}

var OutputNote = make(map[string]Products)

func main() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/api/products", GetProducts).Methods("POST")

	server := &http.Server{
		Addr: ":8585",
		Handler:r,
		ReadTimeout : 10 * time.Second,
		WriteTimeout : 10 * time.Second,
		MaxHeaderBytes : 1 << 20,
	}
	log.Println("Listening http://localhost:8585 ...")
	server.ListenAndServe()
}