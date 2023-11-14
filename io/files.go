package io

import (
	"errors"
	arrays "github.com/rwirdemann/restkit/arrays"
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

// ReadLines reads the contents of filename into an array of strings.
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

func InsertFragment(filename string, before string, fragment string) error {
	lines, err := ReadLines(filename)
	if err != nil {
		return err
	}
	inserted, err := arrays.Insert(lines, before, fragment)
	if err != nil {
		return err
	}

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, line := range inserted {
		_, err := f.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}
