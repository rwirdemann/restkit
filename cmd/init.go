package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var name string
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Creates the project",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Creating project '%s'\n", args[0])
	},
}

func init() {
	initCmd.Flags().StringVar(&name, "name", "", "project name")
	rootCmd.AddCommand(initCmd)
}
