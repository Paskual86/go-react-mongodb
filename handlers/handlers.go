package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/Paskual86/go-react-mongodb.git/middlew"
	"github.com/Paskual86/go-react-mongodb.git/routers"
)

/*Handlers Functions that allow manage the connection*/
func Handlers() {
	router := mux.NewRouter()

	// Armado del registro
	router.HandleFunc("/register", middlew.DatabaseCheck(routers.Register)).Methods("POST")
	router.HandleFunc("/login", middlew.DatabaseCheck(routers.Login)).Methods("POST")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	// El paqueter cors tiene funciones que permite dar privilegios segun el usuario. En este caso con la funcion AllowAll() estamos dejando que cualquiera pueda acceder.
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
