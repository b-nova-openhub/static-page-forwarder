package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string

	rootCmd = &cobra.Command{
		Use:           "stapafor",
		Short:         "stapafor – a service that scraps and exposes a set of web pages",
		Long:          ``,
		Version:       "0.0.0",
		SilenceErrors: true,
		SilenceUsage:  true,
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	serveCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "Path to custom config file. Default is '$HOME/.stapafor/config.yaml'.")
	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(serveCmd)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("/etc/stapafor/")
		viper.AddConfigPath("$HOME/.stapafor/")
		viper.AddConfigPath("../config/")
		viper.AddConfigPath(".")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Config file used for stapafor: ", viper.ConfigFileUsed())
	}
}
