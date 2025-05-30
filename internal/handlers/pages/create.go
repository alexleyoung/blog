package pages

import (
	"net/http"

	"github.com/alexleyoung/blog/ui/components"
)

func Yap(w http.ResponseWriter, r *http.Request) {
	component := components.Create()
	component.Render(r.Context(), w)
}
