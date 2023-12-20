package adapter

import (
	"github.com/rwirdemann/restkit/io"
)

type Env struct {
}

func (e Env) RKPort() (int, error) {
	return io.RKPort()
}

func (e Env) GoPath() (string, error) {
	return io.GoPath()
}
