package add

import (
	"fmt"
	"github.com/rwirdemann/restkit/io"
	"log"
	"os"
	"text/template"
	"unicode"
)

func Execute(resourceName string) error {
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

	return nil
}

func capitalize(str string) string {
	runes := []rune(str)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
