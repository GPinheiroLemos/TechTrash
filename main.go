package main

import (
	"log"
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
	router.Put("/lixeira", controllers.EditLixeira)
	router.Delete("/lixeira", controllers.DeleteLixeira)
	router.Get("/loglixeira", controllers.GetLog)
	router.Post("/loglixeira", controllers.PostLog)
	router.Post("/cadastrarusuario", user.NewUser)
	router.Post("/autenticarusuario", user.AuthUser)
	log.Print("Listenning on port 8000")
	http.ListenAndServe(":8000", router)
}
