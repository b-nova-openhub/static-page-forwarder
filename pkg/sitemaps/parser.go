package sitemaps

import (
	sitemap "github.com/oxffaa/gopher-parse-sitemap"
	"log"
)

func Parse(url string) []string {
	result := make([]string, 0, 0)
	err := sitemap.ParseIndexFromFile(url, func(e IndexEntry) error {
		result = append(result, e.GetLocation())
		return nil
	})
	if err != nil {
		log.Fatalln(err)
	}
	return result
}
