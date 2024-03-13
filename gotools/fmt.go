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
	log.Printf("cmd:    gofmt -s -w %s", filesystem.Pwd())
	cmd := fmt.Sprintf("gofmt -d -w %s", filesystem.Pwd())
	if _, err := exec.Command("bash", "-c", cmd).Output(); err != nil {
		log.Fatalln(err)
		return err
	}

	return nil
}
