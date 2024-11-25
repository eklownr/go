package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/eklownr/pretty"
)

func Persistence(n int) int {
	count := 0
	digits := strings.Split(strconv.Itoa(n), "")
	intSlice := make([]int, len(digits))
	for i, d := range digits {
		intSlice[i], _ = strconv.Atoi(d)
	}
	fmt.Println(intSlice) // [1, 2, 3, 4]

	// // slice := []int{}
	// for i := n; i > 10; i %= 10 {
	// 	count++
	// }
	return count
}

func main() {
	p("39 gives %d", Persistence(39))
	p("999 gives %d", Persistence(999))
	p("4 gives %d", Persistence(4))
	p("25 gives %d", Persistence(25))

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
