package main

import (
	"log"
	"net/http"

	"github.com/alexleyoung/blog/internal/firebase"
	"github.com/alexleyoung/blog/internal/handlers"
	"github.com/alexleyoung/blog/internal/handlers/pages"
)

func main() {
	firebase.InitFirebase()

	PORT := "8080"
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", pages.Home)
	mux.HandleFunc("GET /blab/{slug}", pages.Post)
	mux.HandleFunc("GET /yap", pages.Create)
	mux.HandleFunc("GET /edit", pages.Edit)

	mux.HandleFunc("POST /blab", handlers.PostPost)
	mux.HandleFunc("PATCH /blab", handlers.PatchPost)

	log.Println("Server starting on port :" + PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, mux))
}
