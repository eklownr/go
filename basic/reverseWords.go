package main

import (
	"fmt"
	"strconv"
	"strings"
)

func ReverseWords2(str string) string {
	strSplit := strings.Split(str, " ")
	var res []string
	for i := len(strSplit) - 1; i >= 0; i-- {
		res = append(res, strSplit[i])
	}
	return strings.Join(res, " ")
}

func ReverseWords(str string) string {
	var rw string
	words := strings.Split(str, " ")
	for i, _ := range words {
		i++
		rw += words[len(words)-i]

		if len(words)-i != 0 {
			rw += " "
		}
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

var StringToNumber = strconv.Atoi

func main() {
	w := ReverseWords("hello world here i come")
	fmt.Println(w)
	rr := reverseRunes("world")
	fmt.Println(rr)
	rw := ReverseWords2("sex laxar i en laxask")
	fmt.Println(rw)

	stn, err := StringToNumber("1234")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("strToNr: %v", stn)
	}
}
