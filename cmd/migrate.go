package cmd

import (
	"github.com/rwirdemann/restkit/cli"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(migrateCmd)
}

var migrateCmd = &cobra.Command{
	Use:   "migrate <SQL_FILE>",
	Short: "Migrates database",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := cli.Migrate(args[0]); err != nil {
			return err
		}
		return nil
	},
}
