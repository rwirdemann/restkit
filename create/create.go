package create

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
	"text/template"
)

func Execute(name string) {
	root := viper.GetString("RESTKIT_ROOT")
	if len(root) == 0 {
		fmt.Fprint(os.Stderr, "env RESTKIT_ROOT not set")
		os.Exit(1)
	}

	if _, err := os.Stat(root); errors.Is(err, os.ErrNotExist) {
		fmt.Fprintf(os.Stderr, "Root directory %s does not exist", root)
		os.Exit(1)
	}

	if !strings.HasSuffix(root, string(os.PathSeparator)) {
		root = fmt.Sprintf("%s%s", root, string(os.PathSeparator))
	}

	log.Printf("RESTKIT_ROOT: %s\n", root)
	path := root + name
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		log.Printf("Creating project '%s'...\n", path)
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	} else {
		log.Printf("Project '%s' exists\n", path)
	}

	temp, err := template.ParseGlob("templates/*")
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		Project string
	}{
		Project: name,
	}
	gomod, _ := os.Create(fmt.Sprintf("%s/go.mod", path))
	defer gomod.Close()
	err = temp.ExecuteTemplate(gomod, "go.mod.txt", data)
	if err != nil {
		log.Fatalln(err)
	}

	gomain, _ := os.Create(fmt.Sprintf("%s/main.go", path))
	defer gomain.Close()
	err = temp.ExecuteTemplate(gomain, "main.go.txt", data)
	if err != nil {
		log.Fatalln(err)
	}
}
