package goColor

import (
	"github.com/fatih/color"
)

// replace for fmt.Printf(string, multiple values)
func Pl(s string, value ...interface{}) {
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
