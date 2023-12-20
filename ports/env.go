package ports

type Env interface {
	RKPort() (int, error)
	GoPath() (string, error)
}
