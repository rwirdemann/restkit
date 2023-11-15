package cmd

import (
	"github.com/rwirdemann/restkit/adapter"
	"github.com/rwirdemann/restkit/io"
	"github.com/rwirdemann/restkit/ports"
	"github.com/spf13/cobra"
	"log"
)

var fileSystem ports.FileSystem
var env ports.Env

func init() {
	env = adapter.Env{}

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
	err = io.CreateDirectoryIfNotExits(path)
	if err != nil {
		return err
	}

	data := struct {
		Project string
	}{
		Project: name,
	}
	err = io.Create("go.mod.txt", "go.mod", path, data)
	if err != nil {
		return err
	}
	return io.Create("main.go.txt", "main.go", path, data)
}
