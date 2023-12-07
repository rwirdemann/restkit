package io

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strconv"
	"strings"
)

// RKRoot returns the root directory specified by the environment variable
// RESTKIT_ROOT. This is the directory where new RESTKit projects are created.
// The directory name is always prefixed by os.PathSeparator. The function
// returns an error if the directory doesn't exist.
func RKRoot() (string, error) {
	root := viper.GetString("RESTKIT_ROOT")
	if len(root) == 0 {
		return "", fmt.Errorf("env %s not set", "RESTKIT_ROOT")
	}

	if _, err := os.Stat(root); errors.Is(err, os.ErrNotExist) {
		return "", fmt.Errorf("restkit root directory %s does not exist", root)
	}

	if !strings.HasSuffix(root, string(os.PathSeparator)) {
		root = fmt.Sprintf("%s%s", root, string(os.PathSeparator))
	}

	return root, nil
}

func RKTemplatePath() (string, error) {
	p := viper.GetString("RESTKIT_TEMPLATES")
	if len(p) == 0 {
		return "", fmt.Errorf("env %s not set", "RESTKIT_TEMPLATES")
	}
	if !strings.HasSuffix(p, string(os.PathSeparator)) {
		p = fmt.Sprintf("%s%s", p, string(os.PathSeparator))
	}
	return p, nil
}

func RKPort() (int, error) {
	p := viper.GetString("RESTKIT_PORT")
	if len(p) == 0 {
		return -1, fmt.Errorf("env %s not set", "RESTKIT_PORT")
	}

	if i, err := strconv.Atoi(p); err != nil {
		return -1, errors.New("value of RESTKIT_PORT must be numeric")
	} else {
		return i, nil
	}
}
