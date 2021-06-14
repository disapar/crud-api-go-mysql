package main

import (
	"log"
	"net/http"
	"rest-api-go/commons"
	"rest-api-go/routes"

	"github.com/gorilla/mux"
)

func main() {
	commons.Migrate()
	router := mux.NewRouter()
	routes.SetPersonaRoutes(router)

	server := http.Server{
		Addr:    ":9000",
		Handler: router,
	}

	log.Println("Servidor ejecutandose sobre el puerto 9000")
	log.Panicln(server.ListenAndServe())

}
