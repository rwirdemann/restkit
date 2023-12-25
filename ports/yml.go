package ports

type Config struct {
	Module string `yaml:"module"`
}

type Yml interface {
	ReadConfig() (Config, error)
}
