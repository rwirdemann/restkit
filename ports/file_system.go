package ports

import "os"

type FileSystem interface {
	CreateDir(path string) error
	CreateFile(path string) (*os.File, error)
	Remove(path string) error
	Exists(path string) bool
	Pwd() string
	Base(path string) string
}
