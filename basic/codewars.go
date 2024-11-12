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

func Encode(str string) (encodeStr string) {
	for _, c := range str {
		replace := false
		// if val, ok := codeMap[c]; ok {
		// 	fmt.Printf("val: %v, Ok: %v \n", val, ok)
		// }
		for k, v := range codeMap {
			if c == v {
				encodeStr += string(k)
				replace = true
			}
		}
		if !replace {
			encodeStr += string(c)
		}
	}

	return
}

var codeMap = map[rune]rune{
	'G': 'A', 'g': 'a', 'D': 'E', 'd': 'e', 'R': 'Y', 'r': 'y',
	'P': 'O', 'p': 'o', 'L': 'U', 'l': 'u', 'K': 'I', 'k': 'i',
	'A': 'G', 'a': 'g', 'E': 'D', 'e': 'd', 'Y': 'R', 'y': 'r',
	'O': 'P', 'o': 'p', 'U': 'L', 'u': 'l', 'I': 'K', 'i': 'k',
}

// func Encode(str string) (res string) {
// 	for _, v := range str {
// 		val, ok := codeMap[v]
// 		if ok {
// 			res += string(val)
// 		} else {
// 			res += string(v)
// 		}
// 	}
// 	return res
// }

func Decode(str string) string {
	return Encode(str)
}

func ToJadenCase(str string) string {
	return strings.Title(str)
}

func initName(name string) string {
	name = strings.ToUpper(name)
	words := strings.Split(name, " ")
	var initial []string
	for _, val := range words {
		initial = append(initial, string(val[0]))
	}
	return initial[0] + "." + initial[1]
}

func GetMiddle(s string) string {
	strLen := len(s)
	if strLen%2 == 0 {
		return string(s[strLen/2-1]) + string(s[strLen/2])
	}
	return string(s[strLen/2])
}

func HowMuchILoveYou(i int) string {
	var phrase = []string{
		"I love you",
		"a little",
		"a lot",
		"passionately",
		"madly",
		"not at all",
	}
	for i > 6 {
		i = i - 6
	}
	return phrase[i-1]
}
func GetGrade(a, b, c int) rune {
	average := (a + b + c) / 3
	if average >= 90 {
		return 'A'
	}
	if average >= 80 {
		return 'B'
	}
	if average >= 70 {
		return 'C'
	}
	if average >= 60 {
		return 'D'
	}
	return 'F'
}

func Past(h, m, s int) int {
	return (h*60*60 + m*60 + s) * 1000
}

func TwoToOne(s1 string, s2 string) string {
	s3 := s1 + s2
	return s3
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

	fmt.Println(initName("Richard Eklöw"))

	fmt.Println(GetMiddle("Richar"))
	fmt.Println(GetMiddle("Richard"))
	fmt.Println(GetMiddle("R"))
	fmt.Println(GetMiddle("char"))

	fmt.Printf("\nHowMuchILoveYou: %s\n", HowMuchILoveYou(6))
	fmt.Printf("HowMuchILoveYou: %s\n\n", HowMuchILoveYou(7))

	fmt.Printf("0, 1, 1 => 61000: %v\n", Past(0, 1, 1))
}
