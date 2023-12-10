package cmd

import (
	"fmt"
	"log"
	"os"
	"unicode"

	"github.com/rwirdemann/restkit/gotools"
	"github.com/spf13/cobra"
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

	if err := createHttpHandler(resourceName); err != nil {
		return err
	}

	if err := createDomainObject(resourceName); err != nil {
		return err
	}

	if err := createPorts(resourceName); err != nil {
		return err
	}

	if err := gotools.Fmt(); err != nil {
		return err
	}

	return nil
}

func createHttpHandler(resourceName string) error {
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
	if err := createFromTemplate(fmt.Sprintf("%s_handler.go", pluralize(resourceName)), httpDir, "resource_handler.go.txt", data); err != nil {
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
	check := fmt.Sprintf("%sAdapter := http2.New%sHandler()", pluralize(resourceName), pluralize(capitalize(resourceName)))
	if contains, _ := template.Contains("main.go", check); contains {
		log.Printf("insert: %s...already there\n", "http handler")
	} else {
		log.Printf("insert: %s...ok\n", "http handler")
		fragment := fmt.Sprintf("%sAdapter := http2.New%sHandler()\n"+
			"\trouter.HandleFunc(\"/%s\", %sAdapter.GetAll()).Methods(\"GET\")\n",
			pluralize(resourceName), pluralize(capitalize(resourceName)),
			pluralize(resourceName), pluralize(resourceName))
		if err := template.InsertFragment("main.go", "log.Println(\"starting http service on port 8080...\")", fragment); err != nil {
			return err
		}
	}

	return nil
}

func createDomainObject(resourceName string) error {
	// Create application dir if not exist
	if err := createDirIfNotExists("application"); err != nil {
		return err
	}

	// Create domain dir if not exist
	appDir := fmt.Sprintf("%s%c%s", "application", os.PathSeparator, "domain")
	if err := createDirIfNotExists(appDir); err != nil {
		return err
	}

	// Create domain object for resource representation
	data := struct {
		Resource string
	}{
		Resource: capitalize(resourceName),
	}
	if err := createFromTemplate(fmt.Sprintf("%s.go", resourceName), appDir, "resource.go.txt", data); err != nil {
		return err
	}

	return nil
}

func createPorts(resourceName string) error {
	// Create ports dir if not exists
	if err := createDirIfNotExists("ports"); err != nil {
		return err
	}

	// Create in dir if not exist
	inDir := fmt.Sprintf("%s%c%s", "ports", os.PathSeparator, "in")
	if err := createDirIfNotExists(inDir); err != nil {
		return err
	}

	data := struct {
		Resource string
	}{
		Resource: capitalize(resourceName),
	}
	if err := createFromTemplate(fmt.Sprintf("%s_service.go", pluralize(resourceName)), inDir, "in_port.go.txt", data); err != nil {
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

func pluralize(str string) string {
	return str + "s"
}
