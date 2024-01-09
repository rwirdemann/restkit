package cmd

import (
	"fmt"
	"testing"

	"github.com/rwirdemann/restkit/mocks/github.com/rwirdemann/restkit/ports"

	ports2 "github.com/rwirdemann/restkit/ports"
)

var mockEnv *ports.MockEnv
var mockFileSystem *ports.MockFileSystem
var mockTemplate *ports.MockTemplate
var mockYml *ports.MockYml

func TestAddResource(t *testing.T) {
	createMocks(t)
	testAddResource(false)
}

func TestForceAddResource(t *testing.T) {
	createMocks(t)
	testAddResource(true)
}

func testAddResource(f bool) {
	force = f
	mockFileSystem.EXPECT().Exists(".restkit.yml").Return(true)
	mockFileSystem.EXPECT().Exists("context").Return(false)
	mockFileSystem.EXPECT().CreateDir("context").Return(nil)
	mockFileSystem.EXPECT().Exists("context/http").Return(false)
	mockFileSystem.EXPECT().CreateDir("context/http").Return(nil)

	data := struct {
		Resource string
	}{
		Resource: "Book",
	}
	if f {
		mockFileSystem.EXPECT().Exists("context/http/books_handler.go").Return(true)
		mockFileSystem.EXPECT().Remove("context/http/books_handler.go").Return(nil)
	} else {
		mockFileSystem.EXPECT().Exists("context/http/books_handler.go").Return(false)
	}
	mockTemplate.EXPECT().Create("resource_handler.go.txt", "books_handler.go", "context/http", data).Return(nil)

	mockFileSystem.EXPECT().Pwd().Return("github.com/rwirdemann/bookstore")
	mockFileSystem.EXPECT().Base("github.com/rwirdemann/bookstore").Return("bookstore")
	mockTemplate.EXPECT().Contains("main.go", "http2 \"github.com/rwirdemann/bookstore/context/http\"").Return(false, nil)
	mockTemplate.EXPECT().InsertFragment("main.go",
		"\"net/http\"",
		"http2 \"github.com/rwirdemann/bookstore/context/http\"").Return(nil)

	mockTemplate.EXPECT().Contains("main.go", "booksAdapter := http2.NewBooksHandler()").Return(false, nil)
	mockTemplate.EXPECT().InsertFragment("main.go",
		"log.Printf(\"starting http service on port %d...\", c.Port)",
		"booksAdapter := http2.NewBooksHandler()\n"+
			"\trouter.HandleFunc(\"/books\", booksAdapter.GetAll()).Methods(\"GET\")\n").Return(nil)

	mockFileSystem.EXPECT().Exists("application").Return(false)
	mockFileSystem.EXPECT().CreateDir("application").Return(nil)
	mockFileSystem.EXPECT().Exists("application/domain").Return(false)
	mockFileSystem.EXPECT().CreateDir("application/domain").Return(nil)

	if f {
		mockFileSystem.EXPECT().Exists("application/domain/book.go").Return(true)
		mockFileSystem.EXPECT().Remove("application/domain/book.go").Return(nil)
	} else {
		mockFileSystem.EXPECT().Exists("application/domain/book.go").Return(false)
	}
	mockTemplate.EXPECT().Create("resource.go.txt", "book.go", "application/domain", data).Return(nil)

	portData := struct {
		Resource string
		Project  string
	}{
		Resource: "Book",
		Project:  "bookstore",
	}
	mockFileSystem.EXPECT().Exists("ports").Return(false)
	mockFileSystem.EXPECT().CreateDir("ports").Return(nil)
	mockFileSystem.EXPECT().Exists("ports/in").Return(false)
	mockFileSystem.EXPECT().CreateDir("ports/in").Return(nil)
	mockFileSystem.EXPECT().Exists("ports/out").Return(false)
	mockFileSystem.EXPECT().CreateDir("ports/out").Return(nil)

	// Create service port
	expectCreatePortFiles(f, "books_service.go", "in_port.go.txt", "books_service.go", "ports/in", portData)

	// Create repository port
	expectCreatePortFiles(f, "books_repository.go", "repository_out_port.go.txt", "books_repository.go", "ports/out", portData)

	// Create service
	c := ports2.Config{
		Module: "github.com/rwirdemann/bookstore",
		Port:   8080,
	}
	mockYml.EXPECT().ReadConfig().Return(c, nil)
	serviceData := struct {
		Resource          string
		ResourceLowerCaps string
		Module            string
	}{
		Resource:          "Book",
		ResourceLowerCaps: "book",
		Module:            "github.com/rwirdemann/bookstore",
	}
	mockFileSystem.EXPECT().Exists("application").Return(false)
	mockFileSystem.EXPECT().CreateDir("application").Return(nil)
	mockFileSystem.EXPECT().Exists("application/services").Return(false)
	mockFileSystem.EXPECT().CreateDir("application/services").Return(nil)
	mockFileSystem.EXPECT().Exists("application/services/books.go").Return(false)
	mockTemplate.EXPECT().Create("service.go.txt", "books.go", "application/services", serviceData).Return(nil)

	mockYml.EXPECT().ReadConfig().Return(ports2.Config{}, nil)

	_ = add("book")
}

func expectCreatePortFiles(f bool, portName string, templ string, out string, outPath string, portData struct {
	Resource string
	Project  string
}) {
	portPath := fmt.Sprintf("%s/%s", outPath, portName)
	if f {
		mockFileSystem.EXPECT().Exists(portPath).Return(true)
		mockFileSystem.EXPECT().Remove(portPath).Return(nil)
	} else {
		mockFileSystem.EXPECT().Exists(portPath).Return(false)
	}
	mockTemplate.EXPECT().Create(templ, out, outPath, portData).Return(nil)
}

func createMocks(t *testing.T) {
	mockEnv = ports.NewMockEnv(t)
	env = mockEnv
	mockFileSystem = ports.NewMockFileSystem(t)
	fileSystem = mockFileSystem
	mockTemplate = ports.NewMockTemplate(t)
	template = mockTemplate
	mockYml = ports.NewMockYml(t)
	yml = mockYml
}
