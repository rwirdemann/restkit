package adapter

import (
	"errors"
	"github.com/rwirdemann/restkit/io"
	"strings"
)

type Env struct {
}

func (e Env) RKRoot() (string, error) {
	return io.RKRoot()
}

func (e Env) RKPort() (int, error) {
	return io.RKPort()
}

func (e Env) RKModule() (string, error) {
	var err error
	var root string
	root, err = e.RKRoot()
	if err != nil {
		return "", err
	}

	if !strings.Contains(root, "src/") {
		return "", errors.New("RKRoot contains no src directory")
	}
	return strings.TrimSuffix(strings.Split(root, "src/")[1], "/"), nil
}
