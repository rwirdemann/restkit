package adapter

import (
	"fmt"
	"github.com/rwirdemann/restkit/arrays"
	io2 "github.com/rwirdemann/restkit/io"
	"io"
	"os"
	"strings"
	"text/template"
)

type Template struct {
}

// ReadLines reads the contents of filename into an array of strings.
func ReadLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return []string{}, err
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		return []string{}, err
	}
	lines := strings.Split(string(data), "\n")
	return lines, nil
}

func (t Template) InsertFragment(filename string, before string, fragment string) error {
	lines, err := ReadLines(filename)
	if err != nil {
		return err
	}
	inserted, err := arrays.Insert(lines, before, fragment)
	if err != nil {
		return err
	}

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, line := range inserted {
		_, err := f.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}

func (t Template) Create(templ string, out string, path string, data interface{}) error {
	templatePath, err := io2.RKTemplatePath()
	if err != nil {
		return err
	}

	temp, err := template.ParseGlob(fmt.Sprintf("%s/*", templatePath))
	if err != nil {
		return err
	}

	gomod, err := os.Create(fmt.Sprintf("%s/%s", path, out))
	defer gomod.Close()
	if err != nil {
		return err
	}
	return temp.ExecuteTemplate(gomod, templ, data)
}
