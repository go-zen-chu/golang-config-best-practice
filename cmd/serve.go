package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serve called")
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	// bind command line args to viper key
	serveCmd.PersistentFlags().Int("port", 9090, "Server port")
	viper.BindPFlag("server.port", serveCmd.PersistentFlags().Lookup("port"))
	serveCmd.PersistentFlags().String("github-user", "user", "Github user")
	viper.BindPFlag("github.user", serveCmd.PersistentFlags().Lookup("github-user"))
	serveCmd.PersistentFlags().String("github-secret", "secret", "Github user's secret")
	viper.BindPFlag("github.secret", serveCmd.PersistentFlags().Lookup("github-secret"))
}
