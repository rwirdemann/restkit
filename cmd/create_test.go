package cmd

import (
	"fmt"
	"testing"

	"github.com/rwirdemann/restkit/mocks/github.com/rwirdemann/restkit/ports"
)

func TestCreateProjectDirectory(t *testing.T) {
	mockEnv := ports.NewMockEnv(t)
	env = mockEnv
	mockFileSystem := ports.NewMockFileSystem(t)
	fileSystem = mockFileSystem
	mockTemplate := ports.NewMockTemplate(t)
	template = mockTemplate
	mockTime := ports.NewMockTime(t)
	time = mockTime

	path := "/go/src/github.com/rwirdemann/bookstore"
	data := struct {
		Project          string
		Port             int
		Module           string
		DatabaseDriver   string
		DatabaseName     string
		DatabaseUser     string
		DatabasePassword string
	}{
		Project:        "bookstore",
		Port:           8080,
		Module:         "github.com/rwirdemann/bookstore",
		DatabaseDriver: "postgres",
		DatabaseName:   "bookstore",
		DatabaseUser:   "ralf",
	}

	mockFileSystem.EXPECT().Exists("/go/src/github.com/rwirdemann/bookstore").Return(false)
	mockFileSystem.EXPECT().CreateDir("/go/src/github.com/rwirdemann/bookstore").Return(nil)
	mockFileSystem.EXPECT().Exists("/go/src/github.com/rwirdemann/bookstore/migrations").Return(false)
	mockFileSystem.EXPECT().CreateDir("/go/src/github.com/rwirdemann/bookstore/migrations").Return(nil)
	mockFileSystem.EXPECT().Exists("/go/src/github.com/rwirdemann/bookstore/.restkit.yml").Return(false)
	mockFileSystem.EXPECT().AssertCreated("/go/src/github.com/rwirdemann/bookstore/.restkit.yml")
	mockTemplate.EXPECT().Create("restkit.yml.txt", ".restkit.yml", path, data).Return(nil)

	mockFileSystem.EXPECT().Exists("/go/src/github.com/rwirdemann/bookstore/go.mod").Return(false)
	mockTemplate.EXPECT().Create("go.mod.txt", "go.mod", path, data).Return(nil)

	mockFileSystem.EXPECT().Exists("/go/src/github.com/rwirdemann/bookstore/main.go").Return(false)
	mockTemplate.EXPECT().Create("main.go.txt", "main.go", path, data).Return(nil)

	mockTime.EXPECT().TS().Return("20241212121408")
	migrationsPath := fmt.Sprintf("%s/migrations", path)
	mockFileSystem.EXPECT().Exists("/go/src/github.com/rwirdemann/bookstore/migrations/20241212121408_create_database.sql").Return(false)
	mockTemplate.EXPECT().Create("create_database.sql.txt", "20241212121408_create_database.sql", migrationsPath, data).Return(nil)

	create("github.com/rwirdemann/bookstore", "/go/src/github.com/rwirdemann/bookstore", 8080)
}
