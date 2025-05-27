package pages

import (
	"net/http"

	"github.com/alexleyoung/blog/types"
	"github.com/alexleyoung/blog/ui/components"
)

func Yap(w http.ResponseWriter, r *http.Request) {
	component := components.Yap()
	component.Render(r.Context(), w)
}
