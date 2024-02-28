package gotools

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuild(t *testing.T) {
	f := FragmentBuilder{}
	f.Append("hello")
	assert.Equal(t, "hello", f.Build("book"))
	f.Append("world")
	assert.Equal(t, "hello\nworld", f.Build("book"))
}

func TestPlaceholders(t *testing.T) {
	f := FragmentBuilder{}
	f.Append("hello%r")
	assert.Equal(t, "hellobook", f.Build("book"))
	f.Append("world%R")
	assert.Equal(t, "hellobook\nworldBook", f.Build("book"))
}
