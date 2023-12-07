package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"unicode"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	addCmd.AddCommand(resourceCmd)
}

var resourceCmd = &cobra.Command{
	Use:   "resource",
	Short: "creates a resource",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := add(args[0]); err != nil {
			return err
		}
		return nil
	},
}

func add(resourceName string) error {

	// Check if current directory is a RESTkit's project root
	if !fileSystem.Exists(".restkit") {
		return fmt.Errorf("current directory contains no .restkit")
	}

	// Create context dir if not exists
	if err := createDirIfNotExists("context"); err != nil {
		return err
	}

	// Create http dir if not exist
	httpDir := fmt.Sprintf("%s%c%s", "context", os.PathSeparator, "http")
	if err := createDirIfNotExists(httpDir); err != nil {
		return err
	}

	// Create resource handler file
	data := struct {
		Resource string
	}{
		Resource: capitalize(resourceName),
	}
	if err := createFromTemplate(fmt.Sprintf("%s_handler.go", resourceName), httpDir, "resource_handler.go.txt", data); err != nil {
		return err
	}

	// Insert adapter import statement into main file
	projectName := fileSystem.Base(fileSystem.Pwd())
	f := fmt.Sprintf("http2 \"github.com/rwirdemann/%s/context/http\"", projectName)
	if contains, _ := template.Contains("main.go", f); contains {
		log.Printf("insert: %s...already there\n", "import")
	} else {
		log.Printf("insert: %s...ok\n", "import")
		if err := template.InsertFragment("main.go", "\"net/http\"", f); err != nil {
			return err
		}
	}

	// Insert create adapter into main file
	check := fmt.Sprintf("%sAdapter := http2.New%sHandler()", resourceName, capitalize(resourceName))
	if contains, _ := template.Contains("main.go", check); contains {
		log.Printf("insert: %s...already there\n", "http handler")
	} else {
		log.Printf("insert: %s...ok\n", "http handler")
		fragment := fmt.Sprintf("%sAdapter := http2.New%sHandler()\n"+
			"\trouter.HandleFunc(\"/%ss\", %sAdapter.GetAll()).Methods(\"GET\")\n", resourceName, capitalize(resourceName), resourceName, resourceName)
		if err := template.InsertFragment("main.go", "log.Println(\"starting http service on port 8080...\")", fragment); err != nil {
			return err
		}
	}

	// Create domain dir if not exist
	if err := createDirIfNotExists("domain"); err != nil {
		return err
	}

	// Create domain object for resource representation
	if err := createFromTemplate(fmt.Sprintf("%s.go", resourceName), "domain", "resource.go.txt", data); err != nil {
		return err
	}

	// Run go fmt
	root := viper.GetString("RESTKIT_ROOT")
	if len(root) == 0 {
		return fmt.Errorf("env %s not set", "RESTKIT_ROOT")
	}
	cmd := fmt.Sprintf("go fmt %s", fmt.Sprintf("%s/%s", root, projectName))
	if _, err := exec.Command("bash", "-c", cmd).Output(); err != nil {
		return err
	}

	return nil
}

func createFromTemplate(filename, path, tmpl string, data interface{}) error {
	fn := fmt.Sprintf("%s/%s", path, filename)
	if fileSystem.Exists(fn) {
		log.Printf("create: %s...exists\n", fn)
	} else {
		log.Printf("create: %s...ok\n", path)
		if err := template.Create(tmpl, filename, path, data); err != nil {
			return err
		}
	}
	return nil
}

func createDirIfNotExists(dir string) error {
	if !fileSystem.Exists(dir) {
		if err := fileSystem.CreateDir(dir); err != nil {
			return err
		}
	}
	return nil
}

func capitalize(str string) string {
	runes := []rune(str)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
