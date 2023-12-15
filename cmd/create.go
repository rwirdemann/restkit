package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	createCmd.Flags().BoolVarP(&force, "force", "f", false, "override existing project")
	rootCmd.AddCommand(createCmd)
}

var force = false
var createCmd = &cobra.Command{
	Use:   "create module",
	Short: "Creates the project",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		module := args[0]
		if err := validateModule(module); err != nil {
			return err
		}
		restkitPort, err := env.RKPort()
		if err != nil {
			return err
		}

		goPath, err := env.GoPath()
		if err != nil {
			return err
		}
		projectRoot := fmt.Sprintf("%s%c%s%c%s", goPath, os.PathSeparator, "src", os.PathSeparator, module)
		if force && fileSystem.Exists(projectRoot) {
			if err := remove(projectRoot); err != nil {
				return err
			}
		}

		if err := create(module, projectRoot, restkitPort); err != nil {
			return err
		}

		return nil
	},
}

func remove(projectRoot string) error {
	if _, err := os.Stat(projectRoot); err == nil {
		if err := os.RemoveAll(projectRoot); err != nil {
			return err
		}
		log.Printf("remove: %s...ok\n", projectRoot)
	}

	return nil
}

func create(module string, projectRoot string, port int) error {
	if err := createDirIfNotExist(projectRoot); err != nil {
		return err
	}

	restkit := fmt.Sprintf("%s/.restkit", projectRoot)
	if err := createFileIfNotExist(restkit); err != nil {
		return err
	}

	projectName, err := projectName(module)
	if err != nil {
		return err
	}
	data := struct {
		Project string
		Port    int
		Module  string
	}{
		Project: projectName,
		Port:    port,
		Module:  module,
	}

	if err := createTemplateIfNotExists(projectRoot, data, "go.mod.txt", "go.mod"); err != nil {
		return err
	}

	if err := createTemplateIfNotExists(projectRoot, data, "main.go.txt", "main.go"); err != nil {
		return err
	}

	return nil
}

func createDirIfNotExist(path string) error {
	if fileSystem.Exists(path) {
		log.Printf("create: %s exists\n", path)
	} else {
		log.Printf("create: %s...ok\n", path)
		if err := fileSystem.CreateDir(path); err != nil {
			return err
		}
	}
	return nil
}

func createFileIfNotExist(path string) error {
	if fileSystem.Exists(path) {
		log.Printf("create: %s exists\n", path)
	} else {
		log.Printf("create: %s...ok\n", path)
		if _, err := fileSystem.CreateFile(path); err != nil {
			return err
		}
	}
	return nil
}

func createTemplateIfNotExists(path string, data interface{}, tmpl string, out string) error {
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
