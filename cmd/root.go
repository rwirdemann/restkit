package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var version = "0.0.2"
var rootCmd = &cobra.Command{
	Use:     "rk",
	Version: version,
	Short:   "restkit - a simple CLI to generate rest apis",
	Long: `restkit is a super fancy CLI 
   
One can use restkit to generate customized, postgres backed rest apis in minutes`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
