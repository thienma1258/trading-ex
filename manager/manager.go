package manager

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"trading/config"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "neshr-trading",
	Short: "Use for calculate and trading ",
	Long:  `CLI for starting calculate and trading with profit`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

var checkCmd = &cobra.Command{
	Use:   "run",
	Short: " execute cmd",
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("running trading success")
	},
}

func init() {

	logLevel, err := log.ParseLevel(config.Default.App.LogLevel)
	if err != nil {
		log.WithError(err).Fatal("Failed to parse log level.")
	}

	log.SetLevel(logLevel)
}

func InitCommands() {
	rootCmd.AddCommand(checkCmd)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
