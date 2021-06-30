package main

import (
	"b-nova-openhub/stapafor/pkg/config"
	"b-nova-openhub/stapafor/pkg/rest"
)

func main() {
	config.PersistConfig()
	rest.HandleRequests()
}
