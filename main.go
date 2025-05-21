package main

import (
	"context"
	"os"

	"github.com/alexleyoung/blog/templates"
)

func main() {
	component := templates.Home()
	component.Render(context.Background(), os.Stdout)
}
