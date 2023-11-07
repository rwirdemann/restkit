package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
)

var name string
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Creates the project",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		root := viper.GetString("RESTKIT_ROOT")
		if len(root) == 0 {
			fmt.Fprint(os.Stderr, "env RESTKIT_ROOT not set")
			os.Exit(1)
		}

		if _, err := os.Stat(root); errors.Is(err, os.ErrNotExist) {
			fmt.Fprintf(os.Stderr, "Root directory %s does not exist", root)
			os.Exit(1)
		}

		if !strings.HasSuffix(root, string(os.PathSeparator)) {
			root = fmt.Sprintf("%s%s", root, string(os.PathSeparator))
		}

		log.Printf("RESTKIT_ROOT: %s\n", root)
		path := root + args[0]
		if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
			log.Printf("Preating project '%s'...\n", path)
			err := os.Mkdir(path, os.ModePerm)
			if err != nil {
				log.Println(err)
			}
		} else {
			log.Printf("Project '%s' exists\n", path)
		}
	},
}

func init() {
	initCmd.Flags().StringVar(&name, "name", "", "project name")
	rootCmd.AddCommand(initCmd)
}
