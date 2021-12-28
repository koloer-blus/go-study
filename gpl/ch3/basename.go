package main

import (
	"fmt"
	"strings"
)

func baseName1(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func baseName(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[dot:]
	}
	return s
}

func main() {
	fmt.Print(baseName1("c/d.go"))
	fmt.Print(baseName("a/n.c"))
}
