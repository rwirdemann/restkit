package io

import (
	"log"
	"os"
)

func Remove(name string) error {
	root, err := RKRoot()
	if err != nil {
		return err
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
