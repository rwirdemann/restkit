package io

import (
	"errors"
	"io"
	"log"
	"os"
	"strings"
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

func ReadLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return []string{}, err
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		return []string{}, err
	}
	lines := strings.Split(string(data), "\n")
	return lines, nil
}
