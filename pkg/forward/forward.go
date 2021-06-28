package forward

import "b-nova-openhub/stapafor/pkg/parser"

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
		p.Title = ""
		p.Permalink = ""
		p.Author = ""
		p.Tags = ""
		p.Categories = ""
		p.PublishDate = ""
		p.Description = ""
		p.ShowComments = ""
		p.IsPublished = ""
		p.Body = cp
		pages = append(pages, p)
		parser.ParseByTag(cp, "og:title")
	}

	ForwardedPages = pages

	Forwarded = new(Forwarding)
	Forwarded.Success = true
	Forwarded.Errors = make([]string, 0)
	return Forwarded
}
