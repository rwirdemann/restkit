package adapter

import (
	"fmt"
	"github.com/rwirdemann/restkit/io"
	"os"
	"text/template"
)

type Template struct {
}

func (t Template) Create(templ string, out string, path string, data interface{}) error {
	templatePath, err := io.RKTemplatePath()
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
