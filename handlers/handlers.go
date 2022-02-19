package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Handlers Functions that allow manage the connection*/
func Handlers() {
	router := mux.NewRouter()

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	// El paqueter cors tiene funciones que permite dar privilegios segun el usuario. En este caso con la funcion AllowAll() estamos dejando que cualquiera pueda acceder.
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
