package pages

import (
	"net/http"

	"github.com/alexleyoung/blog/internal/firebase"
	"github.com/alexleyoung/blog/ui/components"
)

func Home(w http.ResponseWriter, r *http.Request) {
	posts, err := firebase.GetAllPosts()
	if err != nil {
		http.Error(w, "error fetching postingtons whoops", http.StatusBadRequest)
		return
	}

	component := components.Blobs(posts)
	component.Render(r.Context(), w)
}
