package cmd

import (
	"github.com/rwirdemann/restkit/adapter"
	"github.com/rwirdemann/restkit/ports"
)

var fileSystem ports.FileSystem
var env ports.Env
var template ports.Template

func init() {
	env = adapter.Env{}
	fileSystem = adapter.FileSystem{}
	template = adapter.Template{}
}
