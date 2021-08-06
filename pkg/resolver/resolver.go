package resolver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GetContentHtml(url string) []string {
	content := make([]string, 0, 0)
	contentUrls := getResolvedUrls(url)
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

func getResolvedUrls(url string) []string {
	get, getErr := http.Get(url)
	if getErr != nil {
		log.Fatal(getErr)
	}
	body, readErr := ioutil.ReadAll(get.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	var resolvedUrls []string
	jsonErr := json.Unmarshal(body, &resolvedUrls)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return resolvedUrls
}
