package handlers

import (
	"net/http"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/alexleyoung/blog/internal/firebase"
	"github.com/alexleyoung/blog/internal/utils"
	"github.com/alexleyoung/blog/types"
)

func PatchPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "error parsing form unlucky", http.StatusBadRequest)
		return
	}

	form := r.PostForm
	title := form.Get("title")
	body := form.Get("body")
	if title == "" || body == "" {
		http.Error(w, "can't have blank haha fail.", http.StatusBadRequest)
		return
	}
	slug := utils.GetSlugFromTitle(title)

	iter := firebase.FirestoreClient.Collection("posts").Where("title", "==", title).Documents(r.Context())
	doc, err := iter.Next()
	if err != nil {
		http.Error(w, "trouble finding post... awkward haha", http.StatusInternalServerError)
		return
	}

	_, err = doc.Ref.Update(r.Context(), []firestore.Update{
		{Path: "Title", Value: title},
		{Path: "Slug", Value: slug},
		{Path: "Body", Value: body},
		{Path: "UpdatedAt", Value: firestore.ServerTimestamp},
	})

	if err != nil {
		http.Error(w, "failed to update rip", http.StatusInternalServerError)
		return
	}

	w.Header().Set("HX-Redirect", "/blab/"+slug)
	w.Write([]byte("updated doc!"))
}

func PostPost(w http.ResponseWriter, r *http.Request) {
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

	err = firebase.CreatePost(r.Context(), post)
	if err != nil {
		http.Error(w, "error creating post gg", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("yap created successfully!"))
}
