package ports

type FileSystem interface {
	CreateDir(path string) error
	Exists(path string) bool
}
