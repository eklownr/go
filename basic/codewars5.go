package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/eklownr/pretty"
)

func Persistence2(n int) int {
	steps := 0
	for n >= 10 {
		m := 1
		for n > 0 {
			m *= n % 10
			n /= 10
		}
		n = m
		steps++
	}
	return steps
}

func Persistence(n int) int {
	count := 0
	prod := 1

	for n >= 10 {
		strOfiInt := strings.Split(strconv.Itoa(n), "")
		intSlice := make([]int, len(strOfiInt))
		for i, d := range strOfiInt {
			intSlice[i], _ = strconv.Atoi(d)
		}
		//fmt.Println(intSlice) // TEST
		for _, v := range intSlice {
			prod *= v
		}
		count++
		n = prod
		prod = 1
		//p("prod: %v", prod) // TEST
	}
	return count
}

func Gimme(array [3]int) int {
	switch {
	case array[0] > array[1] && array[0] < array[2]:
		return 0
	case array[0] > array[2] && array[0] < array[1]:
		return 0
	case array[1] > array[0] && array[1] < array[2]:
		return 1
	case array[1] > array[2] && array[1] < array[0]:
		return 1
	}
	return 2
}

func ToAlternatingCase(str string) string {
	runes := []rune(str)
	s := ""
	for _, r := range runes {
		if unicode.IsUpper(r) {
			s += string(unicode.ToLower(r))
		} else {
			s += string(unicode.ToUpper(r))
		}
	}
	return s
}

func main() {
	p(ToAlternatingCase("aA bB HELLO hello"))
	// p("39 gives %d expect: %d", Persistence(39), 3)
	// p("999 gives %d expect: %d", Persistence(999), 4)
	// p("4 gives %d expect: %d", Persistence(4), 0)
	// p("25 gives %d expect: %d", Persistence(25), 2)

	// pln("***********************", "red")
	// p("Hello, World! ( 1234 ) 1 + 1 = 2")
	// pln("***********************", "blue")
	// pln("***********************", "green")
	// pln("***********************", "cyan")
	// pln("***********************", "magenta")
	// pln("***********************", "yellow")
	// pln("***********************", "white")
	// pln("***********************", "orange")
}

func pln(s string, color string) {
	pretty.PrintColor(s, color)
}
func p(s string, value ...interface{}) {
	fmt.Print(pretty.Pl(s, value...))
}
