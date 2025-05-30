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
		http.NotFound(w, r)
		return
	}

	component := components.Home(posts)
	component.Render(r.Context(), w)
}
