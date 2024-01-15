package adapter

import (
	"errors"
	"log"
	"os"
	"path/filepath"
)

type FileSystem struct {
}

func (f FileSystem) AssertCreated(path string) {
	if !f.Exists(path) {
		log.Fatalf("assert: %s...false\n", path)
	}
	log.Printf("assert: %s...true\n", path)
}

func (f FileSystem) CreateFile(name string) (*os.File, error) {
	return os.Create(name)
}

func (f FileSystem) CreateDir(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func (f FileSystem) Exists(path string) bool {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}

func (f FileSystem) Pwd() string {
	if pwd, err := os.Getwd(); err == nil {
		return pwd
	}
	return ""
}

func (f FileSystem) Base(path string) string {
	return filepath.Base(path)
}

func (f FileSystem) Remove(path string) error {
	return os.Remove(path)
}
