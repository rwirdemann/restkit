package cmd

import (
	"fmt"
	"github.com/rwirdemann/restkit/adapter"
	"log"
	"os"
	"os/user"

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

		if err := create(module, projectRoot, 8080); err != nil {
			return err
		}

		log.Printf("your project '%s' has been created successfully.", module)
		log.Println("### NEXT STEPS ###")
		log.Printf("cd %s\n", projectRoot)
		log.Println("go mod tidy")
		log.Println("go run main.go")
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

	migrationsRoot := fmt.Sprintf("%s/%s", projectRoot, "migrations")
	if err := createDirIfNotExist(migrationsRoot); err != nil {
		return err
	}

	projectName, err := ProjectName(module)
	if err != nil {
		return err
	}

	u, err := user.Current()
	if err != nil {
		return err
	}

	data := struct {
		Project          string
		Port             int
		Module           string
		DatabaseDriver   string
		DatabaseName     string
		DatabaseUser     string
		DatabasePassword string
	}{
		Project:        projectName,
		Port:           port,
		Module:         module,
		DatabaseDriver: "postgres",
		DatabaseName:   projectName,
		DatabaseUser:   u.Username,
	}

	if err := createTemplateIfNotExists(projectRoot, data, "restkit.yml.txt", ".restkit.yml"); err != nil {
		return err
	}

	assertCreated(projectRoot, ".restkit.yml")

	if err := createTemplateIfNotExists(projectRoot, data, "go.mod.txt", "go.mod"); err != nil {
		return err
	}

	if err := createTemplateIfNotExists(projectRoot, data, "main.go.txt", "main.go"); err != nil {
		return err
	}

	fmt.Printf("TIME: %s\n", adapter.Time{}.TS())

	out := fmt.Sprintf("%s_create_database.sql", time.TS())
	if err := createTemplateIfNotExists(migrationsRoot, data, "create_database.sql.txt", out); err != nil {
		return err
	}

	return nil
}

func assertCreated(root string, s string) {
	path := fmt.Sprintf("%s/%s", root, s)
	fileSystem.AssertCreated(path)
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
