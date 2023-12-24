package adapter

import (
	"fmt"
	"github.com/spf13/viper"
)

type Env struct {
}

func (e Env) GoPath() (string, error) {
	p := viper.GetString("GOPATH")
	if len(p) == 0 {
		return "", fmt.Errorf("env %s not set", "GOPATH")
	}
	return p, nil
}
