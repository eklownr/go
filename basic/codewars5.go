package main

import (
	"github.com/eklownr/pretty"
)

func main() {
	pretty.Pl(FindUniq([]float32{1, 1, 1, 2, 1, 1}))
}
