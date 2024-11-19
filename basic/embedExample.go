package main

import (
	_ "embed"
	"fmt"
)

var (
	//go:embed data.txt
	data string
)

func main() {
	fmt.Println(data)

}
