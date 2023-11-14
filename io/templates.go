package io

import (
	"fmt"
	"os"
	"text/template"
)

func Create(templ string, out string, path string, data interface{}) error {
	temp, err := template.ParseGlob("templates/*")
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
