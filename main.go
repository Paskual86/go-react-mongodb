package main

import (
	"log"

	"github.com/Paskual86/go-react-mongodb.git/db"
	"github.com/Paskual86/go-react-mongodb.git/handlers"
)

func main() {

	if db.CheckConnection() == 0 {
		log.Fatal("Without DB connection")
		return
	}

	handlers.Handlers()
}
