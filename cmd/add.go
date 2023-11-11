package cmd

import (
	"github.com/rwirdemann/restkit/add"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds resource",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := add.Execute(args[0]); err != nil {
			log.Panicf("Fatal error %s", err)
		}
	},
}
