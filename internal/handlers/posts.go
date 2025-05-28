package handlers

import (
	"net/http"
	"time"

	"github.com/alexleyoung/blog/internal/firebase"
	"github.com/alexleyoung/blog/internal/utils"
	"github.com/alexleyoung/blog/types"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "error parsing form unlucky", http.StatusBadRequest)
		return
	}

	form := r.PostForm

	post := types.Post{
		Title:     form.Get("title"),
		Slug:      utils.GetSlugFromTitle(form.Get("title")),
		Body:      form.Get("body"),
		Tags:      []string{},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = firebase.CreatePost(post)
	if err != nil {
		http.Error(w, "error creating post gg", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("post created successfully!"))
}
