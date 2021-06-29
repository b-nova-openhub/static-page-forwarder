package url_converter

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"static-page-exposer/pkg/util"
	"strings"
)

type XMLQuery struct {
	Loc string `xml:",chardata"`
}

var l XMLQuery

func GetAllLocUrls() []string {
	return getLoc(GetHTMLPage())
}

func GetHTMLPage() []byte {
	url := "https://b-nova.com/sitemaps/sitemap.xml"
	resp, err := http.Get(url)
	util.HandleError(err)
	defer resp.Body.Close()
	html, err := ioutil.ReadAll(resp.Body)
	util.HandleError(err)
	return html
}

func getLoc(html []byte) []string {
	var locUrls []string

	decoder := xml.NewDecoder(strings.NewReader(string(html)))

	for {
		token, _ := decoder.Token()

		if token == nil {
			break
		}

		switch Element := token.(type) {
		case xml.StartElement:
			if Element.Name.Local == "loc" {
				err := decoder.DecodeElement(&l, &Element)
				util.HandleError(err)
				locUrls = append(locUrls, l.Loc)
			}
		}
	}
	return locUrls
}
