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
	mockFileSystem.EXPECT().Exists("adapter").Return(false)
	mockFileSystem.EXPECT().CreateDir("adapter").Return(nil)
	mockFileSystem.EXPECT().Exists("adapter/http").Return(false)
	mockFileSystem.EXPECT().CreateDir("adapter/http").Return(nil)
	data := struct {
		Resource string
	}{
		Resource: "Book",
	}

	mockFileSystem.EXPECT().Exists("adapter/http/book_handler.go").Return(false)
	mockTemplate.EXPECT().Create("resource_handler.go.txt", "book_handler.go", "adapter/http", data).Return(nil)

	mockFileSystem.EXPECT().Pwd().Return("github.com/rwirdemann/bookstore")
	mockFileSystem.EXPECT().Base("github.com/rwirdemann/bookstore").Return("bookstore")
	mockTemplate.EXPECT().Contains("main.go", "http2 \"github.com/rwirdemann/bookstore/adapter/http\"").Return(false, nil)
	mockTemplate.EXPECT().InsertFragment("main.go",
		"\"net/http\"",
		"http2 \"github.com/rwirdemann/bookstore/adapter/http\"").Return(nil)

	mockTemplate.EXPECT().Contains("main.go", "bookAdapter := http2.NewBookHandler()").Return(false, nil)
	mockTemplate.EXPECT().InsertFragment("main.go",
		"log.Println(\"starting http service on port 8080...\")",
		"bookAdapter := http2.NewBookHandler()\n"+
			"\trouter.HandleFunc(\"/books\", bookAdapter.GetAll()).Methods(\"GET\")\n").Return(nil)

	add("book")
}