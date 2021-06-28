package rest

import (
	"b-nova-openhub/stapafor/pkg/forward"
	"b-nova-openhub/stapafor/pkg/resolver"
	"b-nova-openhub/stapafor/src/github.com/gorilla/mux"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func HandleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/page", getPage).Methods("GET")
	router.HandleFunc("/pages", getPages).Methods("GET")
	router.HandleFunc("/status", getStatus).Methods("GET")
	router.HandleFunc("/forward", getForward).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func getPage(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	pages := forward.ForwardedPages
	page := getPageById(v.Get("id"), pages)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&page)
}

func getPages(w http.ResponseWriter, r *http.Request) {
	pages := forward.ForwardedPages
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pages)
}

func getStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(nil)
}

func getForward(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Forward Request: %+v\n", r)
	html := resolver.GetContentHtml("https://b-nova.com/sitemaps/sitemap.xml")
	forwarded := forward.Forward(html)
	fmt.Printf("Forward Response: %+v\n", forwarded)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(forwarded)
}

func getPageById(id string, pages []forward.StaticPage) *forward.StaticPage {
	var page *forward.StaticPage
	for _, p := range pages {
		if p.Permalink == id {
			page = &p
		}
	}
	return page
}
