package pages

import (
	"net/http"

	"github.com/alexleyoung/blog/ui/components"
)

func Create(w http.ResponseWriter, r *http.Request) {
	component := components.Create()
	component.Render(r.Context(), w)
}
