package cmd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProjectName(t *testing.T) {
	n, _ := ProjectName("github.com/rwirdemann/bookstore")
	assert.Equal(t, "bookstore", n)
}

func TestValidateAttributes(t *testing.T) {
	assert.Nil(t, validateAttributes([]string{"title:string"}))
	assert.Nil(t, validateAttributes([]string{"title:string", "author:string"}))

	assert.Nil(t, validateAttributes([]string{"year:int"}))

	assert.Error(t, validateAttributes([]string{"titlestring"}))
	assert.Error(t, validateAttributes([]string{"title string"}))
	assert.Error(t, validateAttributes([]string{"title"}))
	assert.Error(t, validateAttributes([]string{"title;string"}))
	assert.Error(t, validateAttributes([]string{"title:"}))
}
