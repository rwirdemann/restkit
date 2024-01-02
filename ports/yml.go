package ports

type Config struct {
	Module string `yaml:"module"`
	Port   int    `yaml:"port"`
}

type Yml interface {
	ReadConfig() (Config, error)
}
