package ports

type Env interface {
	GoPath() (string, error)
}
