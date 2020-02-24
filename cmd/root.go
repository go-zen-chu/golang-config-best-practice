package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var configPath string

var rootCmd = &cobra.Command{
	Use:   "golang-config-best-practice",
	Short: "A brief description of your application",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("execute root command: %s", err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	// command line flag for passing config file
	rootCmd.PersistentFlags().StringVar(&configPath, "config", "", "Path to config file")
	rootCmd.Flags().BoolP("verbose", "v", false, "Verbosing output")
}

func initConfig() {
	if configPath != "" {
		if _, err := os.Stat(configPath); err != nil {
			log.Fatalf("stat: %s", err)
		}
		viper.SetConfigFile(configPath)
	}

	// Enable loading MYAPP_SERVER_PORT as `server.port` in viper
	viper.SetEnvPrefix("MYAPP")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
