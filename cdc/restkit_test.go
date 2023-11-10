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
	assert.True(t, exist(projectRoot), "does not exist: "+projectRoot)
}

func exist(file string) bool {
	_, err := os.Stat(file)
	return err == nil
}
