package io

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strconv"
	"strings"
)

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

func GoPath() (string, error) {
	p := viper.GetString("GOPATH")
	if len(p) == 0 {
		return "", fmt.Errorf("env %s not set", "GOPATH")
	}
	return p, nil
}
