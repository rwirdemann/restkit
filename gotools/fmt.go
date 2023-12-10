package gotools

import (
	"fmt"
	"os/exec"

	"github.com/rwirdemann/restkit/adapter"
	"github.com/rwirdemann/restkit/ports"
	"github.com/spf13/viper"
)

var fileSystem ports.FileSystem

func init() {
	fileSystem = adapter.FileSystem{}
}

func Fmt() error {
	root := viper.GetString("RESTKIT_ROOT")
	if len(root) == 0 {
		return fmt.Errorf("env %s not set", "RESTKIT_ROOT")
	}
	projectName := fileSystem.Base(fileSystem.Pwd())
	cmd := fmt.Sprintf("go fmt %s", fmt.Sprintf("%s/%s", root, projectName))
	if _, err := exec.Command("bash", "-c", cmd).Output(); err != nil {
		return err
	}

	return nil
}
