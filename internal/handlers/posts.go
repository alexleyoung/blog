package handlers

import (
	"net/http"

	"github.com/alexleyoung/blog/internal/firebase"
	"github.com/alexleyoung/blog/internal/utils"
	"github.com/alexleyoung/blog/types"
)

func CreatePost(w http.ResponseWriter, r http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	form := r.PostForm

	post := types.Post{
		Title: utils.GetSlugFromTitle(form.Get("title")),
	}

	firebase.CreatePost(post)
}
