package main

import (
	"fmt"

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

	//red := color.New(color.FgRed).PrintfFunc()
	//red("Warning\n")

	//blue := color.New(color.FgHiBlue).PrintfFunc()
	Pl("This message is cyan and a new line ")
	//	Pl("This message is cyan with extra string %s", "...extra...")

	yellow := color.New(color.FgYellow).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	fmt.Printf("This is a %s and this is %s.\n", yellow("warning"), red("error"))

}
