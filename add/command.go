package add

import (
	"errors"
	"fmt"
	"log"
	"os"
	"text/template"
	"unicode"
)

func Execute(resourceName string) error {
	createIfNotExits("adapter")
	httpDir := fmt.Sprintf("%s%c%s", "adapter", os.PathSeparator, "http")
	createIfNotExits(httpDir)

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

func createIfNotExits(name string) {
	if _, err := os.Stat(name); errors.Is(err, os.ErrNotExist) {
		log.Printf("create: %s...ok\n", name)
		err := os.Mkdir(name, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	} else {
		log.Printf("dir '%s' exists\n", name)
	}
}

func capitalize(str string) string {
	runes := []rune(str)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
