package cmd

import (
	"github.com/rwirdemann/restkit/mocks/github.com/rwirdemann/restkit/ports"
	"testing"
)

func TestCreateProjectDirectory(t *testing.T) {
	mockEnv := ports.NewMockEnv(t)
	env = mockEnv
	mockFileSystem := ports.NewMockFileSystem(t)
	fileSystem = mockFileSystem

	mockEnv.EXPECT().RKRoot().Return("/github.com/rwirdemann/", nil)
	mockFileSystem.EXPECT().Exists("/github.com/rwirdemann/bookstore").Return(false)
	mockFileSystem.EXPECT().CreateDir("/github.com/rwirdemann/bookstore").Return(nil)

	create("bookstore")
}
