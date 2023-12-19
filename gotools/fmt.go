package gotools

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/rwirdemann/restkit/adapter"
	"github.com/rwirdemann/restkit/ports"
)

var filesystem ports.FileSystem

func init() {
	filesystem = adapter.FileSystem{}
}

func Fmt() error {
	log.Printf("cmd:    go fmt in %s", filesystem.Pwd())
	cmd := fmt.Sprintf("go fmt %s", filesystem.Pwd())
	if _, err := exec.Command("bash", "-c", cmd).Output(); err != nil {
		return err
	}

	return nil
}
