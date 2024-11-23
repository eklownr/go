package main

import (
	"fmt"
	"regexp"
)

func main() {
	// test regexp
	t := "Hello, World!"
	for i := range t {
		if regexp.MustCompile(`[^a-zA-Z]`).MatchString(string(t[i])) {
			fmt.Printf("Non-letter character found: %c\n", t[i])
		}
	}
}
