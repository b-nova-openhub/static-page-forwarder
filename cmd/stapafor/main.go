package main

import (
	"b-nova-openhub/stapafor/pkg/converter"
	"b-nova-openhub/stapafor/pkg/dynamo"
	"b-nova-openhub/stapafor/pkg/filter"
	"b-nova-openhub/stapafor/pkg/html"
	"log"
)

func main() {
	log.Println("All Tables: ")
	dynamo.ListAllTables()

	log.Println("Item by Key Filter: ")
	filterItem := dynamo.GetItemByKey("config", "filter")
	log.Println(filterItem)

	log.Println("All Urls from Sitemap: ")
	urlsFromSiteMap := converter.GetAllLocUrls()
	log.Println(urlsFromSiteMap)

	log.Println("All allowed Urls: ")
	allowedUrls := filter.FilterUrls(urlsFromSiteMap, filterItem)
	log.Println(allowedUrls)

	html.GetContentForSolr(allowedUrls)
}
