package pages

import (
	"net/http"

	"github.com/alexleyoung/blog/internal/firebase"
	"github.com/alexleyoung/blog/ui/components"
)

func Blab(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")

	posts, err := firebase.GetAllPosts()
	if err != nil {
		http.Error(w, "error fetching postingtons whoops", http.StatusBadRequest)
		http.NotFound(w, r)
		return
	}

	for _, blab := range posts {
		if blab.Slug == slug {
			component := components.Blab(blab)
			component.Render(r.Context(), w)
			return
		}
	}

	http.NotFound(w, r)
}
