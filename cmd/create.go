package cmd

import (
	"fmt"
	"github.com/rwirdemann/restkit/io"
	"github.com/rwirdemann/restkit/remove"
	"github.com/spf13/cobra"
	"log"
	"os"
	"text/template"
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
		if err := remove.Execute(args[0]); err != nil {
			log.Panicf("Fatal error %s", err)
		}

		if err := create(args[0]); err != nil {
			log.Panicf("Fatal error %s", err)
		}
	},
}

func create(name string) error {
	root, err := io.RKRoot()
	if err != nil {
		return err
	}

	path := root + name
	err = io.CreateDirectoryIfNotExits(path)
	if err != nil {
		return err
	}

	temp, err := template.ParseGlob("templates/*")
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		Project string
	}{
		Project: name,
	}
	gomod, _ := os.Create(fmt.Sprintf("%s/go.mod", path))
	defer gomod.Close()
	err = temp.ExecuteTemplate(gomod, "go.mod.txt", data)
	if err != nil {
		log.Fatalln(err)
	}

	gomain, _ := os.Create(fmt.Sprintf("%s/main.go", path))
	defer gomain.Close()
	err = temp.ExecuteTemplate(gomain, "main.go.txt", data)
	if err != nil {
		log.Fatalln(err)
	}
	return nil
}
