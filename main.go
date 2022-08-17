package main

import (
	"log"
	"net/http"
	"os"
	"techTrash/connection"
	"techTrash/controllers"
	"techTrash/user"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	connection.MysqlConnect()
	router := mux.NewRouter()
	router.HandleFunc("/lixeira", controllers.GetLixeira).Methods("GET")
	router.HandleFunc("/loglixeira", controllers.GetLog).Methods("GET")
	router.HandleFunc("/lixeira", controllers.PostLixeira).Methods("POST")
	router.HandleFunc("/loglixeira", controllers.PostLog).Methods("POST")
	router.HandleFunc("/cadastrarusuario", user.NewUser).Methods("POST")
	router.HandleFunc("/autenticarusuario", user.AuthUser).Methods("POST")
	log.Print("Running at port :8000")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}
