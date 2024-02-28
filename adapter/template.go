package adapter

import (
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"

	"github.com/gobuffalo/packr"
	"github.com/rwirdemann/restkit/arrays"
	"github.com/spf13/viper"
)

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
	templatePath, err := templatePath()
	if err != nil {
		return err
	}

	box := packr.NewBox(templatePath)
	s, err := box.FindString(templ)
	if err != nil {
		return err
	}

	tmpl, err := template.New(templ).Parse(s)
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

func templatePath() (string, error) {
	gopath := viper.GetString("GOPATH")
	if len(gopath) == 0 {
		return "", fmt.Errorf("env %s not set", "GOPATH")
	}
	return fmt.Sprintf("%s/src/github.com/rwirdemann/restkit/templates", gopath), nil
}
