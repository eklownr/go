package main

import (
	"embed"
	_ "embed"
	"fmt"
)

var (
	//go:embed data.txt
	data string
)

//go:embed *.go
var pwd embed.FS

func main() {
	fmt.Println(data)

	fmt.Println("All go files in this directory:")

	files, _ := pwd.ReadDir(".")
	for _, file := range files {
		fmt.Println(file.Name())
	}
}
