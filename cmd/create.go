package cmd

import (
	"fmt"
	"github.com/rwirdemann/restkit/io"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	createCmd.Flags().StringVar(&name, "name", "", "project name")
	createCmd.Flags().BoolVarP(&force, "force", "f", false, "override existing project")
	rootCmd.AddCommand(createCmd)
}

var name string
var force = false
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates the project",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Force: ", force)
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

	log.Printf("create: %s...ok\n", fmt.Sprintf("%s/.restkit", path))
	_, err = fileSystem.CreateFile(fmt.Sprintf("%s/.restkit", path))
	if err != nil {
		return err
	}

	data := struct {
		Project string
	}{
		Project: name,
	}

	if err := createIfNotExists(path, data, "go.mod.txt", "go.mod"); err != nil {
		return err
	}

	if err := createIfNotExists(path, data, "main.go.txt", "main.go"); err != nil {
		return err
	}

	return nil
}

func createIfNotExists(path string, data interface{}, tmpl string, out string) error {
	fn := fmt.Sprintf("%s/%s", path, out)
	if fileSystem.Exists(fn) {
		log.Printf("create: %s...exists\n", fn)
	} else {
		if err := template.Create(tmpl, out, path, data); err != nil {
			return err
		}
		log.Printf("create: %s...ok\n", fn)
	}
	return nil
}
