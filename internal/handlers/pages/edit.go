package pages

import (
	"net/http"

	"github.com/alexleyoung/blog/ui/components"
)

func Edit(w http.ResponseWriter, r *http.Request) {
	component := components.Edit()
	component.Render(r.Context(), w)
}
