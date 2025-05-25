package main

import (
	"log"
	"net/http"

	"github.com/alexleyoung/blog/firebase"
	"github.com/alexleyoung/blog/handlers"
)

func main() {
	firebase.InitFirebase()

	PORT := "8080"
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handlers.Home)
	mux.HandleFunc("GET /blab/{slug}", handlers.Blab)
	mux.HandleFunc("GET /yap", handlers.Yap)

	log.Println("Server starting on port :" + PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, mux))
}
