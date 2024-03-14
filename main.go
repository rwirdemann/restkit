package main

import (
	"github.com/rwirdemann/restkit/cmd"
)
import "github.com/spf13/viper"

func main() {
	viper.BindEnv("GOPATH")

	cmd.Execute()
}
