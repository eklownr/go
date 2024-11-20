package main

import (
	"fmt"
	"strings"
)

type MyStr string

func (s MyStr) IsUpperCase() bool {
	return s == MyStr(strings.ToUpper(string(s)))
}
func main() {
	fmt.Println("Hello play ground")
	fmt.Println(MyStr.IsUpperCase("c"))
	fmt.Println(MyStr.IsUpperCase("C"))
	fmt.Println(MyStr.IsUpperCase("ABC"))
}
