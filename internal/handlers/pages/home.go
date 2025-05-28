package pages

import (
	"net/http"

	"github.com/alexleyoung/blog/types"
	"github.com/alexleyoung/blog/ui/components"
)

func Home(w http.ResponseWriter, r *http.Request) {
	component := components.Blobs(types.MockPosts)
	component.Render(r.Context(), w)
}
