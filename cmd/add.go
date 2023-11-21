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
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds resource",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := add(args[0]); err != nil {
			log.Panicf("Fatal error %s", err)
		}
	},
}

func add(resourceName string) error {

	// Check if current directory is a RESTkit's project root
	if !fileSystem.Exists(".restkit") {
		return fmt.Errorf("current directory contains no .restkit")
	}

	if !fileSystem.Exists("adapter") {
		if err := fileSystem.CreateDir("adapter"); err != nil {
			return err
		}
	}

	httpDir := fmt.Sprintf("%s%c%s", "adapter", os.PathSeparator, "http")
	if !fileSystem.Exists(httpDir) {
		if err := fileSystem.CreateDir(httpDir); err != nil {
			return err
		}
	}

	data := struct {
		Resource string
	}{
		Resource: capitalize(resourceName),
	}

	resourceFileName := fmt.Sprintf("%s_handler.go", resourceName)
	log.Printf("create: %s/%s...ok\n", httpDir, resourceFileName)
	err := template.Create("resource_handler.go.txt", resourceFileName, httpDir, data)
	if err != nil {
		log.Fatalln(err)
	}

	// Insert create http handler fragment into main file
	if contains, _ := template.Contains("main.go", "http2 \"github.com/rwirdemann/bookstore/adapter/http\""); contains {
		log.Printf("insert: %s...already there\n", "import")
		err = template.InsertFragment("main.go",
			"\"net/http\"",
			"http2 \"github.com/rwirdemann/bookstore/adapter/http\"")
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		log.Printf("insert: %s...ok\n", "import")
		err = template.InsertFragment("main.go",
			"\"net/http\"",
			"http2 \"github.com/rwirdemann/bookstore/adapter/http\"")
		if err != nil {
			log.Fatalln(err)
		}
	}

	log.Printf("insert: %s...ok\n", "http handler")
	fragment := fmt.Sprintf("%sAdapter := http2.New%sHandler()\n"+
		"\trouter.HandleFunc(\"/%ss\", %sAdapter.GetAll()).Methods(\"GET\")\n", resourceName, capitalize(resourceName), resourceName, resourceName)
	err = template.InsertFragment("main.go",
		"log.Println(\"starting http service on port 8080...\")",
		fragment)
	if err != nil {
		log.Fatalln(err)
	}

	root := viper.GetString("RESTKIT_ROOT")
	if len(root) == 0 {
		return fmt.Errorf("env %s not set", "RESTKIT_ROOT")
	}
	path := fmt.Sprintf("%s/bookstore", root)
	cmd := fmt.Sprintf("go fmt %s", path)

	_, err = exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		log.Fatalln(err)
	}
	return nil
}

func capitalize(str string) string {
	runes := []rune(str)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
