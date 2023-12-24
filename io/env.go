package io

import (
	"fmt"
	"github.com/spf13/viper"
)

func RKTemplatePath() (string, error) {
	gopath := viper.GetString("GOPATH")
	if len(gopath) == 0 {
		return "", fmt.Errorf("env %s not set", "GOPATH")
	}
	return fmt.Sprintf("%s/src/github.com/rwirdemann/restkit/templates", gopath), nil
}
