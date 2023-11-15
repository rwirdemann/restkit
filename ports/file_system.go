package ports

import "os"

type FileSystem interface {
	CreateDir(path string) error
	CreateFile(path string) (*os.File, error)
	Exists(path string) bool
}
