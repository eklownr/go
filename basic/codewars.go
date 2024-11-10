package main

import (
	"fmt"
	"strings"
)

func IsPalindrome(str string) bool {
	str = strings.ReplaceAll(str, " ", "")
	word := strings.ToLower(str)
	runes := []rune(word)
	j := len(runes) - 1
	palindrome := false
	for i := 0; i <= j; i++ {
		if runes[i] == runes[j] {
			palindrome = true
			j--
		} else {
			return false
		}
	}
	return palindrome
}

// // test do not pass "ola salo"
// func IsPalindrome(str string) bool {
// 	str = strings.ToLower(str)
// 	n := len(str)
// 	for i := 0; i < n; i++ {
// 		n -= 1
// 		if str[i] != str[n] {
// 			return false
// 		}
// 	}
// 	return true
// }

func SquareSum(numbers []int) (result int) {
	for _, n := range numbers {
		result += n * n
	}
	return result
}

// n 5 --> [5,4,3,2,1]
func ReverseSeq(n int) (a []int) {
	for i := n; i > 0; i-- {
		a = append(a, i)
	}
	return
}

// // crypt, decrypt GA-DE-RY-PO-LU-KI
// var cipher = map[string]string{"G": "A", "D": "E", "R": "Y", "P": "O", "L": "U", "K": "I"}

// func Encode(str string) (encodeStr string) {
// 	s := strings.ToUpper(str)
// 	for _, c := range s {
// 		replace := false
// 		for k, v := range cipher {
// 			if string(c) == v {
// 				encodeStr += k
// 				replace = true
// 			}
// 		}
// 		if !replace {
// 			encodeStr += string(c)
// 		}
// 	}
// 	return
// }

var codeMap = map[rune]rune{
	'G': 'A', 'g': 'a', 'D': 'E', 'd': 'e', 'R': 'Y', 'r': 'y',
	'P': 'O', 'p': 'o', 'L': 'U', 'l': 'u', 'K': 'I', 'k': 'i',
	'A': 'G', 'a': 'g', 'E': 'D', 'e': 'd', 'Y': 'R', 'y': 'r',
	'O': 'P', 'o': 'p', 'U': 'L', 'u': 'l', 'I': 'K', 'i': 'k',
}

func Encode(str string) (res string) {
	for _, v := range str {
		val, ok := codeMap[v]
		if ok {
			res += string(val)
		} else {
			res += string(v)
		}
	}
	return res
}

func Decode(str string) string {
	return Encode(str)
}

func main() {
	a := "Sallad i Dallas"
	b := "Sirap i Paris"
	c := "Ola Salo"
	d := "hello world"
	e := "a"
	fmt.Printf("'%s' is palindrome: %v\n", a, IsPalindrome(a))
	fmt.Printf("'%s' is palindrome: %v\n", b, IsPalindrome(b))
	fmt.Printf("'%s' is palindrome: %v\n", c, IsPalindrome(c))
	fmt.Printf("'%s' is palindrome: %v\n", d, IsPalindrome(d))
	fmt.Printf("'%s' is palindrome: %v\n", e, IsPalindrome(e))

	// squereASum
	my_list := []int{1, 2, 2}
	fmt.Printf("my:list: %v, squareSum: %v\n\n", my_list, SquareSum(my_list))

	// encode decode
	crypt := Encode(b)
	fmt.Printf("%v encode -> %v\n", b, crypt)
	decrypt := Decode(crypt)
	fmt.Printf("%v decode -> %v\n", crypt, decrypt)

	cat := "Ala has a cat"
	encodeCat := Encode(cat)
	fmt.Printf("%v encode -> %v\n", cat, encodeCat)
	de := Decode(encodeCat)
	fmt.Printf("%v decode -> %v\n", encodeCat, de)
}
