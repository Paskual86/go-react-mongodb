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
	router.HandleFunc("/viewprofile", middlew.DatabaseCheck(middlew.ValidateJWT(routers.ViewProfile))).Methods("GET")
	router.HandleFunc("/updateprofile", middlew.DatabaseCheck(middlew.ValidateJWT(routers.UpdateProfile))).Methods("PUT")
	router.HandleFunc("/inserttweet", middlew.DatabaseCheck(middlew.ValidateJWT(routers.InsertTweet))).Methods("POST")
	router.HandleFunc("/readtweets", middlew.DatabaseCheck(middlew.ValidateJWT(routers.ReadTweets))).Methods("GET")
	router.HandleFunc("/deletetweet", middlew.DatabaseCheck(middlew.ValidateJWT(routers.DeleteTweet))).Methods("DELETE")

	router.HandleFunc("/uploadavatar", middlew.DatabaseCheck(middlew.ValidateJWT(routers.UploadAvatar))).Methods("POST")
	router.HandleFunc("/getavatar", middlew.DatabaseCheck(routers.GetAvatar)).Methods("GET")
	router.HandleFunc("/uploadbanner", middlew.DatabaseCheck(middlew.ValidateJWT(routers.UploadBanner))).Methods("POST")
	router.HandleFunc("/getbanner", middlew.DatabaseCheck(routers.GetBanner)).Methods("GET")
	router.HandleFunc("/addrelation", middlew.DatabaseCheck(middlew.ValidateJWT(routers.InsertRelation))).Methods("POST")
	router.HandleFunc("/deleterelation", middlew.DatabaseCheck(middlew.ValidateJWT(routers.DeleteRelation))).Methods("DELETE")
	router.HandleFunc("/getrelation", middlew.DatabaseCheck(middlew.ValidateJWT(routers.QueryRelation))).Methods("GET")

	router.HandleFunc("/listusers", middlew.DatabaseCheck(middlew.ValidateJWT(routers.ListUsers))).Methods("GET")
	router.HandleFunc("/gettweetsfollowers", middlew.DatabaseCheck(middlew.ValidateJWT(routers.GetTweetsFollowers))).Methods("GET")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	// El paqueter cors tiene funciones que permite dar privilegios segun el usuario. En este caso con la funcion AllowAll() estamos dejando que cualquiera pueda acceder.
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
