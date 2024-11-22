package main

import (
	"github.com/fatih/color"
)

func main() {
	// Create a new color object
	c := color.New(color.FgCyan).Add(color.Bold)
	g := color.New(color.FgGreen)

	// Print colored text
	c.Println("This text is cyan and bold!")
	g.Println("This text is green and NOT bold!")

	// Reset color
	color.Unset()

	red := color.New(color.FgRed).PrintfFunc()
	red("Warning\n")

	pl := color.New(color.FgHiGreen).PrintfFunc()
	p := color.New(color.FgHiBlue).PrintfFunc()
	pl("This message is green new line \n")
	p("This message is blue")
}
