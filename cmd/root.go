package cmd

import (
	"fmt"
	"github.com/rwirdemann/restkit/create"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	createCmd.Flags().StringVar(&name, "name", "", "project name")
	rootCmd.AddCommand(createCmd)
}

var version = "0.0.1"
var rootCmd = &cobra.Command{
	Use:     "restkit",
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

var name string
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates the project",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		create.Execute(args[0])
	},
}
