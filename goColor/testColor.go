package main

import (
	"fmt"
	goColor "goColor/color"
	"regexp"

	"golang.org/x/text/number"
)

// A lasy way to print with color
func pl(format string, a ...interface{}) {
	colorize := goColor.Pl(format, a...)
	fmt.Println(colorize)
}

func ll(text string) {
	colorized := goColor.Colorize(text)
	fmt.Println(colorized)
}

// print purple text
func Purple(s string) {
	fmt.Printf("\033[35m%s\033[0m", s)
}
func pc(text string, color string) {
	goColor.PrintColor(text, color)
}

func main() {
	pl("This is a %s and this is %s.", "warning", "error")

	// print type of numbers
	numbers := number.Decimal("0987654321")
	pl("This is numbers 1234567890. %T", numbers)

	// test to many values
	pl("%v %v %v %v %v %v %v", "a", "b", "c", "d", "e", "f", "g")

	// test regexp
	t := "Hello, World!"
	for i := range t {
		if regexp.MustCompile(`[^a-zA-Z]`).MatchString(string(t[i])) {
			fmt.Printf("Non-letter character found: %c\n", t[i])
		}
	}

	//test colorize
	ll("This is a Cyan text with red error <-- Green numbers: ( 1 123 654.0 ) å ä ö")
	ll("asdfasdf, (1234) () *** <> abcdefghijklmnopqrstuvwxyzåäö")

	Purple("This is a purple text")
}
