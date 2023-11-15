package cmd

import (
	"fmt"
	"github.com/rwirdemann/restkit/io"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	createCmd.Flags().StringVar(&name, "name", "", "project name")
	rootCmd.AddCommand(createCmd)
}

var name string
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates the project",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := io.Remove(args[0]); err != nil {
			log.Panicf("Fatal error %s", err)
		}

		if err := create(args[0]); err != nil {
			log.Panicf("Fatal error %s", err)
		}
	},
}

func create(name string) error {
	root, err := env.RKRoot()
	if err != nil {
		return err
	}

	path := root + name
	if fileSystem.Exists(path) {
		log.Printf("create: directory '%s' exists\n", path)
	} else {
		log.Printf("create: %s...ok\n", path)
		if err := fileSystem.CreateDir(path); err != nil {
			return err
		}
	}

	_, err = fileSystem.CreateFile(fmt.Sprintf("%s/.restkit", path))
	if err != nil {
		return err
	}

	data := struct {
		Project string
	}{
		Project: name,
	}
	err = template.Create("go.mod.txt", "go.mod", path, data)
	if err != nil {
		return err
	}
	return template.Create("main.go.txt", "main.go", path, data)
}
