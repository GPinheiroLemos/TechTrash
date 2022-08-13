package main

import (
	"log"
	"net/http"
	"techTrash/connection"
	"techTrash/handlers"

	"github.com/gorilla/mux"
)

func main() {
	connection.MysqlConnect()
	router := mux.NewRouter()
	router.HandleFunc("/lixeira", handlers.GetLixeira).Methods("GET")
	router.HandleFunc("/loglixeira", handlers.GetLog).Methods("GET")
	router.HandleFunc("/lixeira", handlers.PostLixeira).Methods("POST")
	router.HandleFunc("/loglixeira", handlers.PostLog).Methods("POST")
	log.Print("Running at port :8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
