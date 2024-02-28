package gotools

import (
	"fmt"
	"strings"
	"unicode"
)

type FragmentBuilder struct {
	parts []string
}

func (f *FragmentBuilder) Append(p string) {
	f.parts = append(f.parts, p)
}

func (f FragmentBuilder) Build(resource string) string {
	result := ""
	for _, p := range f.parts {
		if result == "" {
			result = p
		} else {
			result = fmt.Sprintf("%s\n%s", result, p)
		}
	}
	result = strings.ReplaceAll(result, "%r", resource)
	result = strings.ReplaceAll(result, "%R", capitalize(resource))
	return result
}

func capitalize(str string) string {
	runes := []rune(str)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
