package main

import (
	"log"
	"net/http"

	"github.com/alexleyoung/blog/firebase"
	"github.com/alexleyoung/blog/handlers"
)

func main() {
	PORT := "8080"

	firebase.InitFirebase()

	http.HandleFunc("/", handlers.Home)

	log.Println("Listening on port " + PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}
