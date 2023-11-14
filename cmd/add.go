package cmd

import (
	"fmt"
	"github.com/rwirdemann/restkit/io"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/exec"
	"text/template"
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
	// check if current directory is project root

	if err := io.CreateDirectoryIfNotExits("adapter"); err != nil {
		return err
	}
	httpDir := fmt.Sprintf("%s%c%s", "adapter", os.PathSeparator, "http")
	if err := io.CreateDirectoryIfNotExits(httpDir); err != nil {
		return err
	}

	temp, err := template.ParseGlob("../restkit/templates/*")
	if err != nil {
		log.Fatalln(err)
	}
	data := struct {
		Resource string
	}{
		Resource: capitalize(resourceName),
	}
	resourceFileName := fmt.Sprintf("%s/%s_handler.go", httpDir, resourceName)
	resourceFile, _ := os.Create(resourceFileName)
	defer resourceFile.Close()
	log.Printf("create: %s...ok\n", resourceFileName)
	err = temp.ExecuteTemplate(resourceFile, "resource_handler.go.txt", data)
	if err != nil {
		log.Fatalln(err)
	}

	// Insert create http handler fragment into main file
	log.Printf("insert: %s...ok\n", "import")
	err = io.InsertFragment("main.go",
		"\"net/http\"",
		"http2 \"github.com/rwirdemann/bookstore/adapter/http\"")
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("insert: %s...ok\n", "http handler")
	err = io.InsertFragment("main.go",
		"err := http.ListenAndServe(fmt.Sprintf(\":%s\", \"8080\"), router)",
		"bookAdapter := http2.NewBookHandler()\n"+
			"\trouter.HandleFunc(\"/books\", bookAdapter.GetAll()).Methods(\"GET\")\n")
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
