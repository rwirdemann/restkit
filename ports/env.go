package ports

type Env interface {
	RKRoot() (string, error)
}
