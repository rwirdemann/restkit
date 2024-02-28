package ports

type Template interface {
	Create(templ string, out string, path string, data interface{}) error
	Insert(filename string, before string, fragment string) error
	Contains(filename string, fragment string) (bool, error)
}
