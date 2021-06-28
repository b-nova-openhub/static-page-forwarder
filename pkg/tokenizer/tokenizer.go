package tokenizer

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"log"
	"strings"
)

func GetTokenByTag(s string, tag string) string {
	tokenizer := html.NewTokenizer(strings.NewReader(s))
	for {
		tokenType := tokenizer.Next()
		if tokenType == html.ErrorToken {
			err := tokenizer.Err()
			if err == io.EOF {
				break
			}
			log.Fatalf("error tokenizing HTML: %v", tokenizer.Err())
		}

		if tokenType == html.StartTagToken {
			token := tokenizer.Token()
			if tag == token.Data {
				tokenType = tokenizer.Next()
				if tokenType == html.TextToken {
					data := tokenizer.Token().Data
					fmt.Println(data)
					if s != "" {
						return data
					}
				}
			}
		}
	}
	return ""
}
