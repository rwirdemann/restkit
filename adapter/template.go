package adapter

import (
	"embed"
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"

	"github.com/rwirdemann/restkit/arrays"
)

//go:embed templates
var templates embed.FS

type Template struct {
}

func (t Template) Contains(filename string, fragment string) (bool, error) {
	lines, err := ReadLines(filename)
	if err != nil {
		return false, err
	}
	index, _ := arrays.Find(lines, fragment)
	return index > -1, nil
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

func (t Template) Insert(filename string, before string, fragment string) error {
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
	f, err := templates.ReadFile(fmt.Sprintf("templates/%s", templ))
	if err != nil {
		return err
	}

	tmpl, err := template.New(templ).Parse(string(f))
	if err != nil {
		return err
	}

	gomod, err := os.Create(fmt.Sprintf("%s/%s", path, out))
	defer gomod.Close()
	if err != nil {
		return err
	}
	return tmpl.ExecuteTemplate(gomod, templ, data)
}
