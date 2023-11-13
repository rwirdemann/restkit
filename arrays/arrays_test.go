package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsert(t *testing.T) {
	before := "err := http.ListenAndServe(fmt.Sprintf(\":%s\", \"8080\"), router)"
	source := []string{
		"println(\"starting http service for project {{.Project}}...\")",
		"router := mux.NewRouter()",
		before,
	}

	actual, err := Insert(source, before, "createGetHandler()")
	assert.Nil(t, err)
	expected := []string{
		"println(\"starting http service for project {{.Project}}...\")",
		"router := mux.NewRouter()",
		"createGetHandler()",
		before,
	}
	assert.Equal(t, expected, actual)
}
