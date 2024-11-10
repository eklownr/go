package main

import (
	"fmt"
	"strconv"
	"strings"
)

func ReverseWords(str string) string {
	strSplit := strings.Split(str, " ")
	var res []string
	for i := len(strSplit) - 1; i >= 0; i-- {
		res = append(res, strSplit[i])
	}
	return strings.Join(res, " ")
}

func ReverseWordsAndRunes(str string) string {
	var rw string
	words := strings.Split(str, " ")

	for i := len(words) - 1; i >= 0; i-- {
		rw += reverseRunes(words[i]) // reverse runes in the word and reverse all words
		if len(words)-i != 0 {       // add spaces between words
			rw += " "
		}
	}
	return rw
}

func reverseRunes(word string) string {
	var rr string
	for i, _ := range word {
		i++
		rr += string(word[len(word)-i])
	}
	return rr
}

var StringToNumber = strconv.Atoi // use var as a function

func main() {
	w := ReverseWordsAndRunes("hello world here i come")
	fmt.Println(w)
	rr := reverseRunes("world")
	fmt.Println(rr)
	rw := ReverseWords("sex laxar i en laxask")
	fmt.Println(rw)

	stn, err := StringToNumber("1234")
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("strToNr: %v", stn)
	}
}
