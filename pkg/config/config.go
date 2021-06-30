package config

import (
	"b-nova-openhub/stapafor/pkg/util"
	"flag"
	"fmt"
)

type Config struct {
	AppPort          string
	TargetSitemapUrl string
}

var AppConfig *Config

func PersistConfig() {
	port := flag.String("port", "8080", "The port being used for the API. Default port is 8080.")
	sitemap := flag.String("sitemap", "https://b-nova.com/sitemaps/sitemap.xml", "This is the fully qualified URL where the desired sitemap to forward sits.")
	flag.Parse()

	AppConfig = new(Config)
	AppConfig.AppPort = util.DerefString(port)
	AppConfig.TargetSitemapUrl = util.DerefString(sitemap)

	fmt.Printf("App configuration for stapafor: \n%+v\n", AppConfig)
}
