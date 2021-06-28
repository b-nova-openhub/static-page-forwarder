package sitemaps

func GetSitemap(url string) []string {
	return Parse(url)
}
