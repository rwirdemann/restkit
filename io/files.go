package io

import (
	"errors"
	"log"
	"os"
)

func CreateDirectoryIfNotExits(name string) error {
	if _, err := os.Stat(name); errors.Is(err, os.ErrNotExist) {
		log.Printf("create: %s...ok\n", name)
		if err := os.Mkdir(name, os.ModePerm); err != nil {
			return err
		}
	} else {
		log.Printf("create: directory '%s' exists\n", name)
	}
	return nil
}
