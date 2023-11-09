package cmd

import (
	"github.com/rwirdemann/restkit/create"
	"github.com/spf13/cobra"
)

var name string
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates the project",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		create.Execute(args[0])
	},
}

func init() {
	createCmd.Flags().StringVar(&name, "name", "", "project name")
	rootCmd.AddCommand(createCmd)
}
