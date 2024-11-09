package main

import (
	"fmt"
	"strings"
)

func ReverseWords(str string) string {
	var rw string
	words := strings.Split(str, " ")
	for i, _ := range words {
		i++
		rw += words[len(words)-i] + " "
	}
	return rw // reversed words
}

func reverseRunes(word string) string {
	var rr string
	for i, _ := range word {
		//fmt.Printf("i: %v, r: %v \n", i, r)
		i++
		rr += string(word[len(word)-i])

	}
	return rr
}

func main() {
	w := ReverseWords("hello world here I come")
	fmt.Println(w)
	rr := reverseRunes("world")
	fmt.Println(rr)
}
