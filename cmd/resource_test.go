package cmd

import (
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

	domainObjectData := struct {
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
	if f {
		mockFileSystem.EXPECT().Exists("ports/in/books_service.go").Return(true)
		mockFileSystem.EXPECT().Remove("ports/in/books_service.go").Return(nil)
	} else {
		mockFileSystem.EXPECT().Exists("ports/in/books_service.go").Return(false)
	}
	mockTemplate.EXPECT().Create("in_port.go.txt", "books_service.go", "ports/in", domainObjectData).Return(nil)

	serviceObjectData := struct {
		Resource string
	}{
		Resource: "Books",
	}
	mockFileSystem.EXPECT().Exists("application").Return(false)
	mockFileSystem.EXPECT().CreateDir("application").Return(nil)
	mockFileSystem.EXPECT().Exists("application/services").Return(false)
	mockFileSystem.EXPECT().CreateDir("application/services").Return(nil)
	mockFileSystem.EXPECT().Exists("application/services/books.go").Return(false)
	mockTemplate.EXPECT().Create("service.go.txt", "books.go", "application/services", serviceObjectData).Return(nil)

	mockYml.EXPECT().ReadConfig().Return(ports2.Config{}, nil)

	_ = add("book")
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
