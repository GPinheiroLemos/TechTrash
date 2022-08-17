package main

import (
	"net/http"
	"techTrash/connection"
	"techTrash/controllers"
	"techTrash/user"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main() {

	connection.MysqlConnect()
	router := chi.NewRouter()
	router.Use(cors.AllowAll().Handler)
	router.Get("/lixeira", controllers.GetLixeira)
	router.Post("/lixeira", controllers.PostLixeira)
	router.Get("/loglixeira", controllers.GetLog)
	router.Post("/loglixeira", controllers.PostLog)
	router.Post("/cadastrarusuario", user.NewUser)
	router.Post("/autenticarusuario", user.AuthUser)
	http.ListenAndServe(":8000", router)
}
