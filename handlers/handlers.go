package handlers

import (
	"net/http"

	"github.com/alexleyoung/blog/components"
	"github.com/alexleyoung/blog/types"
)

func Home(w http.ResponseWriter, r *http.Request) {
	component := components.Blobs(types.MockPosts)
	component.Render(r.Context(), w)
}

func Blab(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	for _, blab := range types.MockPosts {
		if blab.Slug == slug {
			component := components.Blab(blab)
			component.Render(r.Context(), w)
			return
		}
	}
	http.NotFound(w, r)
}

func Yap(w http.ResponseWriter, r *http.Request) {
	component := components.Yap()
	component.Render(r.Context(), w)
}
