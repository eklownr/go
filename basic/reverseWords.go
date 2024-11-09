package main

import (
  "fmt"
  "strings"
)

func ReverseWords(str string) string {
  var rw  string
  words:= strings.Split(str, " ")
  for i, w := range words {
    fmt.Printf("{i: %v,  w: %v\n", i, w)
    rw += w[len(words)-i]
  }
  return rw // reverse those words
}

func main () {
  w := ReverseWords("hello world here I come")
  fmt.Println(w)
}
