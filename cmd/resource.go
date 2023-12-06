package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/exec"
	"unicode"
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

	// Create adapter dir if not exists
	if err := createDirIfNotExists("adapter"); err != nil {
		return err
	}

	// Create http dir if not exist
	httpDir := fmt.Sprintf("%s%c%s", "adapter", os.PathSeparator, "http")
	if err := createDirIfNotExists(httpDir); err != nil {
		return err
	}

	// Create resource handler file
	data := struct {
		Resource string
	}{
		Resource: capitalize(resourceName),
	}
	resourceHandlerFileName := fmt.Sprintf("%s_handler.go", resourceName)
	path := fmt.Sprintf("%s/%s", httpDir, resourceHandlerFileName)
	if fileSystem.Exists(path) {
		log.Printf("create: %s...exists\n", path)
	} else {
		log.Printf("create: %s...ok\n", path)
		err := template.Create("resource_handler.go.txt", resourceHandlerFileName, httpDir, data)
		if err != nil {
			log.Fatalln(err)
		}
	}

	// Insert adapter import statement into main file
	projectName := fileSystem.Base(fileSystem.Pwd())
	f := fmt.Sprintf("http2 \"github.com/rwirdemann/%s/adapter/http\"", projectName)
	if contains, _ := template.Contains("main.go", f); contains {
		log.Printf("insert: %s...already there\n", "import")
	} else {
		log.Printf("insert: %s...ok\n", "import")
		err := template.InsertFragment("main.go", "\"net/http\"", f)
		if err != nil {
			log.Fatalln(err)
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
		err := template.InsertFragment("main.go",
			"log.Println(\"starting http service on port 8080...\")",
			fragment)
		if err != nil {
			log.Fatalln(err)
		}
	}

	// Create domain dir if not exist
	if err := createDirIfNotExists("domain"); err != nil {
		return err
	}

	// Create domain object for resource representation
	resourceFileName := fmt.Sprintf("%s.go", resourceName)
	path = fmt.Sprintf("%s/%s", "domain", resourceFileName)
	if fileSystem.Exists(path) {
		log.Printf("create: %s...exists\n", path)
	} else {
		log.Printf("create: %s...ok\n", path)
		if err := template.Create("resource.go.txt", resourceFileName, "domain", data); err != nil {
			return err
		}
	}

	// Run go fmt
	root := viper.GetString("RESTKIT_ROOT")
	if len(root) == 0 {
		return fmt.Errorf("env %s not set", "RESTKIT_ROOT")
	}
	path = fmt.Sprintf("%s/%s", root, projectName)
	cmd := fmt.Sprintf("go fmt %s", path)
	if _, err := exec.Command("bash", "-c", cmd).Output(); err != nil {
		return err
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
