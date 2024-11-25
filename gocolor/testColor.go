package main

import (
	"fmt"
	goColor "gocolor/color"

	"golang.org/x/text/number"
)

// A lasy way to fmt.Printf() with color
func pl(format string, a ...interface{}) {
	colorize := goColor.Pl(format, a...)
	fmt.Print(colorize)
}

// print colorized textString
func pc(text string, color string) {
	goColor.PrintColor(text, color)
}
func blue(text string) {
	goColor.PrintColor(text, "blue")
}
func orange(text string) {
	goColor.PrintColor(text, "orange")
}

func main() {
	pl("This is a %s and this. is %s. ", "warning", "error")

	// print type of numbers
	numbers := number.Decimal("0987654321")
	pl("This! is numbers (1234567890). This is numbers ( 1234567890 ). %v som text red!! ", "error")
	pl(" special char: ! !! ? . , + -  _ / #  ( )  {} numbers: %v", numbers)
	pl("1 + 1 = 2; 2+2=4")

	// test many values
	pl("%v %v %v %v %v %v %v", "a", "2", "c", "4", "e", "6", "g")

	//test colorize string
	pc("This is a blue text", "blue")
	pc("This is a red text", "red")
	pc("This is a green text", "green")
	pc("This is a yellow text", "yellow")
	pc("This is a magenta text", "magenta")
	pc("This is a cyan text", "cyan")
	pc("This is a white text", "white")
	pc("This is default color", "")
	pc("This is a orange text", "orange")
	blue("more blue!")
	orange("more orange!")
}
