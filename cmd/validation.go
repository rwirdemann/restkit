package cmd

import (
	"errors"
	"os"
	"strings"
)

func validateModule(module string) error {
	s := strings.Split(module, string(os.PathSeparator))
	if len(s) < 2 {
		return errors.New("module name should be of format 'github.org/user/projectname'")
	}

	return nil
}

func projectName(module string) (string, error) {
	if err := validateModule(module); err != nil {
		return "", err
	}

	s := strings.Split(module, string(os.PathSeparator))
	return s[len(s)-1], nil
}
