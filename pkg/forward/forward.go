package forward

import (
	"github.com/b-nova-openhub/stapafor/pkg/tokenizer"
)

type StaticPage struct {
	Title        string `json:"title"`
	Permalink    string `json:"permalink"`
	Author       string `json:"author"`
	Tags         string `json:"tags"`
	Categories   string `json:"categories"`
	PublishDate  string `json:"publishDate"`
	Description  string `json:"description"`
	ShowComments string `json:"showComments"`
	IsPublished  string `json:"isPublished"`
	Body         string `json:"body"`
}

type Forwarding struct {
	Success bool     `json:"success"`
	Errors  []string `json:"errors"`
}

var ForwardedPages []StaticPage
var Forwarded *Forwarding

func Forward(contentPages []string) *Forwarding {
	pages := make([]StaticPage, 0)

	for _, cp := range contentPages {
		var p StaticPage
		p.Title = tokenizer.GetTokenByTag(cp, "title")
		p.Permalink = tokenizer.GetTokenByTag(cp, "og:url")
		p.Author = ""
		p.Tags = ""
		p.Categories = ""
		p.PublishDate = ""
		p.Description = ""
		p.ShowComments = ""
		p.IsPublished = "true"
		p.Body = cp
		pages = append(pages, p)
	}

	ForwardedPages = pages

	Forwarded = new(Forwarding)
	Forwarded.Success = true
	Forwarded.Errors = make([]string, 0)
	return Forwarded
}
