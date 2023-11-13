package arrays

import (
	"fmt"
	"strings"
)

func Insert(source []string, before string, fragment string) ([]string, error) {
	index, err := find(source, before)
	if err != nil {
		return source, err
	}
	return insert(source, index, fragment), nil
}

func find(a []string, s string) (int, error) {
	for i, v := range a {
		if strings.TrimSpace(v) == s {
			return i, nil
		}
	}
	return -1, fmt.Errorf("string not found in array: %s", s)
}

func insert(a []string, index int, value string) []string {
	if len(a) == index {
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...)
	a[index] = value
	return a
}
