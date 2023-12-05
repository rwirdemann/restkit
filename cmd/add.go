package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	addCmd.PersistentFlags().BoolVarP(&force, "force", "f", false, "override existing artefact")
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "adds an artefact",
}
