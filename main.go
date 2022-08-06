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
	log.Fatal(http.ListenAndServe(":8000", router))
}
