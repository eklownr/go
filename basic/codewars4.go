package main

import (
	"fmt"
	"math"
	"strings"
)

type MyStr string

func (s MyStr) IsUpperCase() bool {
	return s == MyStr(strings.ToUpper(string(s)))
}

func Derive(coefficient, exponent int) string {
	return fmt.Sprintf("%dx^%d", coefficient*exponent, exponent-1)
}

func NearestSq(n int) int {
	root := math.Sqrt(float64(n))
	//	roundRoot := math.Round(math.Sqrt(float64(n)))
	//	fmt.Printf("Root: %v, rouned: %v, squere: %v\n", root, roundRoot, roundRoot*roundRoot)

	if int((root)+1)*(int(root)+1)-n < n-int(root)*int(root) {
		return int((root)+1) * (int(root) + 1)
	}
	return int(root) * int(root)
}

func main() {
	fmt.Println("Hello play ground")
	fmt.Println(MyStr.IsUpperCase("c"))
	fmt.Println(MyStr.IsUpperCase("C"))
	fmt.Println(MyStr.IsUpperCase("ABC"))

	fmt.Println(NearestSq(9))
	fmt.Println(NearestSq(10))
	fmt.Println(NearestSq(8))
	fmt.Println(NearestSq(111))
}
