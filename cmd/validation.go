package cmd

import (
	"errors"
	"fmt"
	"github.com/rwirdemann/restkit/arrays"
	"os"
	"regexp"
	"strings"
)

func validateModule(module string) error {
	s := strings.Split(module, string(os.PathSeparator))
	if len(s) < 2 {
		return errors.New("module name should be of format 'github.org/user/projectname'")
	}

	return nil
}

func validateAttributes(attributes []string) error {
	r, err := regexp.Compile("(^\\w+)(:\\w+)")
	if err != nil {
		return err
	}
	for _, a := range attributes {
		if match := r.MatchString(a); !match {
			return errors.New(fmt.Sprintf("attribute '%s' is of invalid format", a))
		}
		s := strings.Split(a, ":")
		if !isValidType(s[1]) {
			return errors.New(fmt.Sprintf("attribute '%s' is has invalid type", a))
		}
	}
	return nil
}

func isValidType(s string) bool {
	return arrays.Contains([]string{"string", "int"}, s)
}

func ProjectName(module string) (string, error) {
	if err := validateModule(module); err != nil {
		return "", err
	}

	s := strings.Split(module, string(os.PathSeparator))
	return s[len(s)-1], nil
}
