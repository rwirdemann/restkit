package cmd

import (
	"errors"
	"fmt"
	"github.com/rwirdemann/restkit/remove"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
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

		if err := execute(args[0]); err != nil {
			log.Panicf("Fatal error %s", err)
		}
	},
}

func execute(name string) error {
	root := viper.GetString("RESTKIT_ROOT")
	if len(root) == 0 {
		return fmt.Errorf("env %s not set", "RESTKIT_ROOT")
	}

	if _, err := os.Stat(root); errors.Is(err, os.ErrNotExist) {
		fmt.Fprintf(os.Stderr, "Root directory %s does not exist", root)
		os.Exit(1)
	}

	if !strings.HasSuffix(root, string(os.PathSeparator)) {
		root = fmt.Sprintf("%s%s", root, string(os.PathSeparator))
	}

	path := root + name
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		log.Printf("create: %s...ok\n", path)
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	} else {
		log.Printf("project '%s' exists\n", path)
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
