package cmd

import (
	"github.com/rwirdemann/restkit/mocks/github.com/rwirdemann/restkit/ports"
	"testing"
)

func TestCreateProjectDirectory(t *testing.T) {
	mockEnv := ports.NewMockEnv(t)
	env = mockEnv
	mockEnv.EXPECT().RKRoot().Return("/github.com/rwirdemann", nil)

	create("bookstore")
}
