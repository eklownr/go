package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	a := "asdf_asdf-apa"
	b := "åäö_åäö-åäö_ABC"
	c := camelCase(a)
	d := camelCase(b)
	fmt.Printf("*********\nCamelCase of: '%s' and: '%s'\n", a, b)
	fmt.Printf("\n********\n return value of camelCase(%s): \n %v\n********\n", a, c)
	fmt.Printf("\n********\n return value of camelCase(%s): \n %v\n********\n", b, d)
}

func camelCase(s string) string {
	var stringOfRunes string // return value
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		isUppercase := true
		if runes[i] == rune('_') || runes[i] == rune('-') {
			isUppercase = false
		}
		if !isUppercase {
			r := unicode.ToUpper(runes[i+1])
			runes[i+1] = rune(r) // skip one rune, set next to uppercase
			isUppercase = true
			continue
		} else {
			stringOfRunes += string(runes[i]) // add all runes to string
		}
	}
	return stringOfRunes
}

// not working
func makeCamelCase(s string) string {
	words := strings.Split(s, "_")
	var camelCase string
	for _, word := range words {
		if len(camelCase) > 0 {
			camelCase += strings.Title(word)
			continue
		} else {
			camelCase = strings.ToLower(word)
		}
	}
	return camelCase
}
