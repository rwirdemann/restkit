package ports

type Template interface {
	Create(templ string, out string, path string, data interface{}) error
}
