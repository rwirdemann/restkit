package adapter

import (
	"errors"
	"os"
)

type FileSystem struct {
}

func (f FileSystem) CreateDir(path string) error {
	return os.Mkdir(path, os.ModePerm)
}

func (f FileSystem) Exists(path string) bool {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}
