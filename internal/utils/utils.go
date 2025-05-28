package utils

import "strings"

func GetSlugFromTitle(title string) string {
	slug := strings.ToLower(title)
	slug = strings.ReplaceAll(slug, " ", "-")
	return slug
}
