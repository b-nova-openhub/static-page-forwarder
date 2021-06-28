package content

import (
	"fmt"
	"github.com/yterajima/go-sitemap"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func getContentHtml(url string) []string {
	content := make([]string, 0, 0)
	contentUrls := getValidContentUrls(url)
	for _, url := range contentUrls {
		get, getErr := http.Get(url)
		if getErr != nil {
			fmt.Println(getErr)
		}
		html, getErr := ioutil.ReadAll(get.Body)
		closeErr := get.Body.Close()
		if closeErr != nil {
			return nil
		}
		content = append(content, string(html))
	}
	return content
}

func getValidContentUrls(url string) []string {
	content := make([]string, 0, 0)
	urls := parseSitemap(url).URL
	for _, url := range urls {
		if strings.Contains("/content/", url.Loc) {
			content = append(content, url.Loc)
		}
	}
	return content
}

func parseSitemap(url string) sitemap.Sitemap {
	smap, err := sitemap.Get(url, nil)
	if err != nil {
		fmt.Println(err)
	}
	for _, URL := range smap.URL {
		log.Println("Found following URL in sitemap.xml: ", URL.Loc)
	}
	return smap
}
