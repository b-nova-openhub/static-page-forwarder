package html

import (
	"b-nova-openhub/stapafor/pkg/util"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
)

//type StaticPage struct {
//	Title        string `json:"title"`
//	Permalink    string `json:"permalink"`
//	Author       string `json:"author"`
//	Tags         string `json:"tags"`
//	Categories   string `json:"categories"`
//	PublishDate  string `json:"publishDate"`
//	Description  string `json:"description"`
//	ShowComments string `json:"showComments"`
//	IsPublished  string `json:"isPublished"`
//	Body         string `json:"body"`
//}

type StaticPage struct {
	Title        string
	Permalink    string
	Author       string
	Tags         string
	Categories   string
	PublishDate  string
	Description  string
	ShowComments string
	IsPublished  string
	Body         string
}

func GetContentForSolr(urls []string) {
	//streams.
	//	FromArray(urls).
	//	ForEach(func(url interface{}) {
	//		urlValue := url.(string)
	//		response, err := http.Get(urlValue)
	//		util.HandleError(err)
	//		defer response.Body.Close()
	//
	//		meta := GetPageContent(response.Body)
	//
	//	log.Println("----------------Start--------------------")
	//	log.Println(meta.Description)
	//	log.Println(meta.Title)
	//
	//	})
	log.Println(urls[0])
	response, err := http.Get(urls[0])
	util.HandleError(err)
	defer response.Body.Close()

	meta := GetPageContent(response.Body)

	log.Println("----------------Start--------------------")
	log.Println(meta.Title)
	log.Println(meta.Author)
	log.Println(meta.Description)
	log.Println(meta.Tags)
	log.Println(meta.PublishDate)
}

func GetPageContent(response io.Reader) (staticPage StaticPage) {
	z := html.NewTokenizer(response)
	for {
		tt := z.Next()

		switch tt {
		case html.ErrorToken:
			return
		case html.StartTagToken:
			t := z.Token()
			if t.Data == "meta" {
				log.Println("------------------------------------------------")
				log.Println(t.Data)

				title, ok := getMetaTagContentByName(t, "title")
				if ok {
					staticPage.Title = title
				}

				author, ok := getMetaTagContentByName(t, "author")
				if ok {
					staticPage.Author = author
				}

				//user, ok := getMetaTagContentByName(t, "user")
				//if ok {
				//	//staticPage.U = user
				//}

				desc, ok := getMetaTagContentByName(t, "description")
				if ok {
					staticPage.Description = desc
				}

				keywords, ok := getMetaTagContentByName(t, "keywords")
				if ok {
					staticPage.Tags = keywords
				}

				publishDate, ok := getMetaTagContentByName(t, "article:published_time")
				if ok {
					staticPage.PublishDate = publishDate
				}
			}
		}
	}
}

func getMetaTagContentByName(t html.Token, prop string) (content string, ok bool) {
	log.Println("getMetaTagContentByName:")
	log.Println(prop)
	log.Println(t)
	for _, attr := range t.Attr {
		if attr.Key == "property" && attr.Val == prop {
			ok = true
		}

		if attr.Key == "name" && attr.Val == prop {
			ok = true
		}

		if attr.Key == "content" {
			content = attr.Val
		}
	}
	return
}
