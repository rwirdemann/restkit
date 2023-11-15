package main

import "github.com/rwirdemann/restkit/cmd"
import "github.com/spf13/viper"

func main() {
	viper.BindEnv("RESTKIT_ROOT")
	viper.BindEnv("RESTKIT_TEMPLATES")

	cmd.Execute()
}
