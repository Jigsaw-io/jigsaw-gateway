package main

import (
	"fmt"
	// "log"
	"net/http"
	"os"

	"github.com/zeemzo/jigsaw-gateway/api/routes"
	"github.com/gorilla/handlers"
	// "github.com/joho/godotenv"
)

func getPort() string {
	p := os.Getenv("PORT")
	if p != "" {
		return ":" + p
	}
	return ":8000"
}


func main() {

	// getEnvironment()

	port := getPort()
	headersOk := handlers.AllowedHeaders([]string{"Content-Type","Access-Control-Allow-Origin"})
	originsOk := handlers.AllowedOrigins([]string{"https://jigsaw-io.firebaseapp.com","http://localhost:3000","https://jigsaw.cf"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	

	router := routes.NewRouter()
	fmt.Println("Jigsaw Gateway Started @port " + port)
	http.ListenAndServe(port, handlers.CORS(originsOk, headersOk, methodsOk)(router))

}
