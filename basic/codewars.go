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

func main() {
	a := "sallad i dallas"
	b := "sirap i paris"
	c := "ola salo"
	d := "hello world"
	e := "a"
	fmt.Printf("'%s' is palindrome: %v\n", a, IsPalindrome(a))
	fmt.Printf("'%s' is palindrome: %v\n", b, IsPalindrome(b))
	fmt.Printf("'%s' is palindrome: %v\n", c, IsPalindrome(c))
	fmt.Printf("'%s' is palindrome: %v\n", d, IsPalindrome(d))
	fmt.Printf("'%s' is palindrome: %v\n", e, IsPalindrome(e))

	// squereASum
	my_list := []int{1, 2, 2}
	fmt.Printf("my:list: %v, squareSum: %v", my_list, SquareSum(my_list))
}
