package cmd

import (
	"github.com/rwirdemann/restkit/mocks/github.com/rwirdemann/restkit/ports"
	"testing"
)

func TestAddResource(t *testing.T) {
	mockEnv := ports.NewMockEnv(t)
	env = mockEnv
	mockFileSystem := ports.NewMockFileSystem(t)
	fileSystem = mockFileSystem
	mockTemplate := ports.NewMockTemplate(t)
	template = mockTemplate

	mockFileSystem.EXPECT().Exists(".restkit").Return(true)
	data := struct {
		Resource string
	}{
		Resource: "Book",
	}

	mockTemplate.EXPECT().Create("resource_handler.go.txt", "book_handler.go", "adapter/http", data).Return(nil)

	add("book")
}
