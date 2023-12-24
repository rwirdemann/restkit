package io

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"strconv"
)

func RKTemplatePath() (string, error) {
	gopath, err := GoPath()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/src/github.com/rwirdemann/restkit/templates", gopath), nil
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
