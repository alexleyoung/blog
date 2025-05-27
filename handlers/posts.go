package handlers

import (
	"net/http"

	"github.com/alexleyoung/blog/firebase"
	"github.com/alexleyoung/blog/types"
)

func CreatePost(w http.ResponseWriter, r http.Request) {
	post := types.Post{}

	firebase.CreatePost(post)
}
