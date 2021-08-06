package cmd

import (
	"github.com/b-nova-openhub/stapafor/pkg/rest"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	port    string
	siteMap string

	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Serve filter-sitemap-customizer service",
		Long:  ``,
		Run:   serve,
	}
)

func init() {
	serveCmd.PersistentFlags().StringVarP(&port, "port", "p", "8080", "This is the port that the REST endpoint is being served on. Default port if none is being specified is 8080.")
	serveCmd.PersistentFlags().StringVarP(&siteMap, "sitemap", "s", "", "This is the fully qualified url where the target sitemap resides.")
	viper.BindPFlag("port", serveCmd.PersistentFlags().Lookup("port"))
	viper.BindPFlag("sitemap", serveCmd.PersistentFlags().Lookup("sitemap"))
}

func serve(ccmd *cobra.Command, args []string) {
	rest.HandleRequests()
}
