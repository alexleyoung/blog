package pages

import (
	"net/http"

	"github.com/alexleyoung/blog/types"
	"github.com/alexleyoung/blog/ui/components"
)

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
