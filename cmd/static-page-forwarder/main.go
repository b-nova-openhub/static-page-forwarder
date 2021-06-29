package main

import (
	"log"
	"static-page-exposer/pkg/dynamoDb"
	"static-page-exposer/pkg/html"
	"static-page-exposer/pkg/url-converter"
	"static-page-exposer/pkg/url-filter-predicate"
)

func main() {
	log.Println("All Tables: ")
	dynamoDb.ListAllTables()

	log.Println("Item by Key Filter: ")
	filterItem := dynamoDb.GetItemByKey("config", "filter")
	log.Println(filterItem)

	log.Println("All Urls from Sitemap: ")
	urlsFromSiteMap := url_converter.GetAllLocUrls()
	log.Println(urlsFromSiteMap)

	log.Println("All allowed Urls: ")
	allowedUrls := url_filter_predicate.FilterUrls(urlsFromSiteMap, filterItem)
	log.Println(allowedUrls)

	html.GetContentForSolr(allowedUrls)
}
