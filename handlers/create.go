package handlers

import (
	"context"
	"net/http"

	"github.com/alexleyoung/blog/templates"
)

func Create(w http.ResponseWriter, r *http.Request) {
	component := templates.Home()
	component.Render(context.Background(), w)
}
