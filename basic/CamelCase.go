package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("\nCamelCase asdf_asdf-apa: ")
	c := camelCase("asdf_asdf-apa")
	d := camelCase("åäö_åäö-åäö")
	fmt.Printf("\n********\n return value of camelCase(): \n%v\n********\n", c)
	fmt.Printf("\n********\n return value of camelCase(): \n%v\n********\n", d)
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
