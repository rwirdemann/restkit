package cdc

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func init() {
	viper.BindEnv("PROJECT_NAME")
	viper.BindEnv("RESTKIT_ROOT")
}

func TestCreate(t *testing.T) {
	root := viper.GetString("RESTKIT_ROOT")
	name := viper.GetString("PROJECT_NAME")

	projectRoot := fmt.Sprintf("%s/%s", root, name)
	assertExists(t, projectRoot)
	assertExists(t, fmt.Sprintf("%s/%s", projectRoot, "go.mod"))
	assertExists(t, fmt.Sprintf("%s/%s", projectRoot, "main.go"))
}

func assertExists(t *testing.T, file string) {
	assert.True(t, exist(file), "does not exist: "+file)
}

func exist(file string) bool {
	_, err := os.Stat(file)
	return err == nil
}
