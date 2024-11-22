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

	//blue := color.New(color.FgHiBlue).PrintfFunc()
	pl("This message is cyan and a new line ")
	pl("This message is cyan with extra string %s", "...extra...")
}

// replace for fmt.Printf(string, multiple values)
func pl(s string, value ...interface{}) {
	green := color.New(color.FgHiGreen).PrintfFunc()
	cyan := color.New(color.FgHiCyan).PrintfFunc()
	i := len(value)
	switch i {
	case 0:
		cyan(s + "\n")
	case 1:
		cyan(s+"\n", value[0])
	case 2:
		cyan(s+"\n", value[0], value[1])
	case 3:
		cyan(s+"\n", value[0], value[1], value[2])
	case 4:
		cyan(s+"\n", value[0], value[1], value[2], value[3])
	case 5:
		cyan(s+"\n", value[0], value[1], value[2], value[3], value[4])
	case 6:
		cyan(s+"\n", value[0], value[1], value[2], value[3], value[4], value[5])
	default:
		green("Too many values")
	}
}
