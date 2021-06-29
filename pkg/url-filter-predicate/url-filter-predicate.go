package url_filter_predicate

import (
	"github.com/jucardi/go-streams/streams"
	"path"
	"static-page-exposer/pkg/util"
	"strings"
	"unicode/utf8"
)

type Item struct {
	value string
}

func FilterUrls(urls []string, filters []string) []string {
	var allowedUrl bool
	allowedUrls := streams.
		FromArray(urls).
		Filter(func(url interface{}) bool {
			urlValue := url.(string)
			for _, filter := range filters {
				if strings.HasPrefix(filter, "!") {
					_, i := utf8.DecodeRuneInString(filter)
					match, err := path.Match("https://b-nova.com"+filter[i:], urlValue)
					allowedUrl = !match
					util.HandleError(err)
				} else {
					match, err := path.Match("https://b-nova.com"+filter, urlValue)
					allowedUrl = match
					util.HandleError(err)
				}
				if allowedUrl {
					break
				}
			}
			return allowedUrl
		}).
		ToArray().([]string)
	return allowedUrls
}
