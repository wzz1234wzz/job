package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("hello!")
}

func Split(s, sep string) []string {
	var result []string
	i := strings.Index(s, sep)
	for i > -1 {
		result = append(result, s[:i])
		s = s[i+len(sep):]
		i = strings.Index(s, sep)
	}
	return append(result, s)
}

func sum(a, b int) int {
	return a + b
}
