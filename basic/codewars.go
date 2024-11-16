package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
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

// cleaner solution
func encode(str string) (res string) {
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
	s := strings.Split(s1+s2, "")
	m := make(map[string]int)
	uniqueStr := ""
	for i, c := range s {
		m[c] = i
	}
	for key, _ := range m {
		uniqueStr += key
	}
	//fmt.Println(m)
	unique := strings.Split(uniqueStr, "")
	sort.Strings(unique)
	return strings.Join(unique, "")
}
func tto(s1, s2 string) (result string) {
	s := s1 + s2
	for _, char := range strings.Split("abcdefghijklmnopqrstuvwxyz", "") {
		if strings.Contains(s, char) {
			result += char
		}
	}
	return
}

func SumMix(arr []any) (sum int) {
	var StringToNumber = strconv.Atoi
	for i, _ := range arr {
		n, ok := arr[i].(int)
		if !ok {
			num, err := StringToNumber(arr[i].(string))
			if err != nil {
				panic(err)
			}
			sum += num
		}
		sum += n
	}
	return
}

func RepeatStr(rep int, value string) (s string) {
	for i := 0; i < rep; i++ {
		s += value
	}
	return s
}

func FindMultiples(integer, limit int) []int {
	resp := []int{}

	for i := integer; i <= limit; i += integer {
		resp = append(resp, i)
	}

	return resp
}
func Hero(bullets, dragons int) bool {
	return bullets >= 2*dragons
}

// return employed && !vacation
func SetAlarm(employed, vacation bool) bool {
	if employed == true && vacation == false {
		return true
	}
	return false
}

func RemoveChar(word string) string {
	return word[1 : len(word)-1]
}

func ReverseLetters(s string) (res string) {
	// reg expression only letters
	s = regexp.MustCompile(`[^a-zA-Z]+`).ReplaceAllString(s, "")

	for i, char := range s {
		i++
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			res += string(s[len(s)-i])
		}
	}
	return res
}

// out += fmt.Sprintf("%d sheep...", i)
func countSheep(num int) (s string) {
	for i := 1; i <= num; i++ {
		s += strconv.Itoa(i) + " sheep..."
	}
	return
}
func QuarterOf(month int) int {
	if month <= 3 {
		return 1
	}
	if month <= 6 {
		return 2
	}
	if month <= 9 {
		return 3
	}
	return 4
}

// accum("RqaEzty") -> "R-Qq-Aaa-Eeee-Zzzzz-Tttttt-Yyyyyyy"
func Accum(s string) string {
	parts := make([]string, len(s))

	for i := 0; i < len(s); i++ {
		parts[i] = strings.ToUpper(string(s[i])) + strings.Repeat(strings.ToLower(string(s[i])), i)
	}
	return strings.Join(parts, "-")
}

// Expect(BinToDec("101010")).To(Equal(42))
func BinToDec(bin string) int {
	num, _ := strconv.ParseInt(bin, 2, 64)

	return int(num)
}
func FindShort(s string) int {
	shortest := len(s)
	for _, word := range strings.Split(s, " ") {
		if len(word) < shortest {
			shortest = len(word)
		}
	}
	return shortest

}

// return strings.ReplaceAll(word, " ", "")
func NoSpace(s string) string {
	words := strings.Split(s, " ")
	return strings.Join(words, "")
}
func century(year int) int {
	return (year + 99) / 100
}
func HighAndLow(num string) string {
	nums := strings.Split(num, " ")
	min, _ := strconv.Atoi(nums[0])
	max := min
	for _, num := range nums {
		n, _ := strconv.Atoi(num)
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return fmt.Sprintf("%d %d", max, min)
}

func StringToArray(str string) []string {
	return strings.Split(str, " ")
}

func High(s string) string {
	words := strings.Split(s, " ")
	highest := 0
	highestWord := ""

	for _, word := range words {
		sum := 0
		for _, char := range word {
			sum += int(char) - 96
		}
		if sum > highest {
			highest = sum
			highestWord = word
		}
	}
	return highestWord
}

func CalculateYears(years int) (result [3]int) {
	result[0] = years
	result[1] = 15 + 9
	result[2] = 15 + 9
	cat := 4
	dog := 5

	for i := 1; i < years; i++ {
		result[2] += cat
		result[1] += dog
	}
	return
}

// remove all aeiouy
func Disemvowel(comment string) string {
	vowels := "aeiouyAEIOUY"
	for _, v := range vowels {
		comment = strings.ReplaceAll(comment, string(v), "")
	}
	return comment
}

func BouncingBall(h, bounce, window float64) int {
	if h <= 0 || bounce <= 0 || bounce >= 1 || window >= h {
		return -1
	}
	count := 1
	for h*bounce > window {
		h *= bounce
		count += 2
	}
	return count
}

func LongestConsec(strarr []string, k int) (result string) {
	for i := 0; i < len(strarr)-k+1; i++ {
		buffer := strings.Join(strarr[i:i+k], "")
		if len(buffer) > len(result) {
			result = buffer
		}
	}
	return
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

	fmt.Println(GetMiddle("Richar"))
	fmt.Println(GetMiddle("Richard"))
	fmt.Println(GetMiddle("R"))
	fmt.Println(GetMiddle("char"))

	fmt.Printf("\nHowMuchILoveYou: %s\n", HowMuchILoveYou(6))
	fmt.Printf("HowMuchILoveYou: %s\n\n", HowMuchILoveYou(7))

	fmt.Printf("0, 1, 1 => 61000: %v\n", Past(0, 1, 1))

	a = "xyaabbbccccdefww"
	b = "xxxxyyyyabklmopq"
	fmt.Printf("a + b: %v\n", TwoToOne(a, b))
	fmt.Printf("a + b: %v\n", tto(a, b))

	num := []any{9, 3, "7", "3"}
	fmt.Printf("any: %v\n", SumMix(num))

	fmt.Printf("multi values %v\n", FindMultiples(2, 6))
	fmt.Printf("multi values %v\n", FindMultiples(5, 25))

	fmt.Printf("Hero wins over dragon %v\n", Hero(0, 1))

	fmt.Printf("Rever letters 'ultr53o?n': %v\n", ReverseLetters("ultr53o?n"))

	fmt.Printf("%v\n", countSheep(3))

	fmt.Printf("%v\n", Accum("ZpglnRxqenU"))
	fmt.Printf("%v\n", Accum("RqaEzty"))

	fmt.Printf("shotest%v\n", FindShort("123 hej på dej:w"))

	fmt.Printf("HighAndLow(1 2 3 4 5): %v\n", HighAndLow("1 2 3 4 5"))

	fmt.Printf("abc: %v\n", High("abc epa"))

	fmt.Printf("calk years: %v\n", CalculateYears(2))

	fmt.Printf("abcxyz: %v\n", Disemvowel("abcxyz"))

	var arr = []string{"zone", "abigail", "theta", "form", "libe", "zas"}
	fmt.Printf("LongestConsec: %v\n", LongestConsec(arr, 2))

	var arr2 = []string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"}
	fmt.Printf("LongestConsec: %v\n", LongestConsec(arr2, 1))
}
