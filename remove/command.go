package remove

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
)

func Execute(name string) error {
	root := viper.GetString("RESTKIT_ROOT")
	if len(root) == 0 {
		return fmt.Errorf("env %s not set", "RESTKIT_ROOT")
	}

	if !strings.HasSuffix(root, string(os.PathSeparator)) {
		root = fmt.Sprintf("%s%s", root, string(os.PathSeparator))
	}

	path := root + name
	if _, err := os.Stat(path); err == nil {
		if err := os.RemoveAll(path); err != nil {
			return err
		}
		log.Printf("remove: %s...ok\n", path)
	}

	return nil
}
