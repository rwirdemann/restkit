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
	mockTemplate := ports.NewMockTemplate(t)
	template = mockTemplate

	mockEnv.EXPECT().RKRoot().Return("/github.com/rwirdemann/", nil)
	mockFileSystem.EXPECT().Exists("/github.com/rwirdemann/bookstore").Return(false)
	mockFileSystem.EXPECT().CreateDir("/github.com/rwirdemann/bookstore").Return(nil)
	mockFileSystem.EXPECT().CreateFile("/github.com/rwirdemann/bookstore/.restkit").Return(nil, nil)

	path := "/github.com/rwirdemann/bookstore"
	data := struct {
		Project string
	}{
		Project: "bookstore",
	}
	mockFileSystem.EXPECT().Exists("/github.com/rwirdemann/bookstore/go.mod").Return(false)
	mockTemplate.EXPECT().Create("go.mod.txt", "go.mod", path, data).Return(nil)

	mockTemplate.EXPECT().Create("main.go.txt", "main.go", path, data).Return(nil)

	create("bookstore")
}
