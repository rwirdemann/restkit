package cmd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProjectName(t *testing.T) {
	n, _ := projectName("github.com/rwirdemann/bookstore")
	assert.Equal(t, "bookstore", n)
}
