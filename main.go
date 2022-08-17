package main

import (
	"net/http"
	"techTrash/connection"
	"techTrash/controllers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main() {

	connection.MysqlConnect()
	router := chi.NewRouter()
	router.Use(cors.AllowAll().Handler)
	router.Get("/lixeira", controllers.GetLixeira)
	http.ListenAndServe(":8000", router)
}
