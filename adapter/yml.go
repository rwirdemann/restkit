package adapter

import (
	"github.com/rwirdemann/restkit/ports"
	"gopkg.in/yaml.v2"
	"os"
)

type Yml struct{}

func (y Yml) ReadConfig() (ports.Config, error) {
	data, err := os.ReadFile(".restkit.yml")
	if err != nil {
		return ports.Config{}, err
	}

	var c ports.Config
	if err := yaml.Unmarshal(data, &c); err != nil {
		return ports.Config{}, err
	}

	return c, nil
}
