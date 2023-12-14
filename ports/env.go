package ports

type Env interface {
	RKRoot() (string, error)
	RKPort() (int, error)
	RKModule() (string, error)
	GoPath() (string, error)
}
