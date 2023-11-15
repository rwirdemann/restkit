package adapter

import "github.com/rwirdemann/restkit/io"

type Env struct {
}

func (e Env) RKRoot() (string, error) {
	return io.RKRoot()
}
