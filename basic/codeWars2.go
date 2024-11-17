package main

import (
	"fmt"
	"strconv"
	"strings"
)

func maxSumSubarray(arr []int, k int) int {
	windowSum := 0
	maxSum := 0
	windowStart := 0

	for windowEnd := 0; windowEnd < len(arr); windowEnd++ {
		// Add the current element to the window sum
		windowSum += arr[windowEnd]

		// If the window size exceeds k, slide the window
		if windowEnd >= k-1 {
			maxSum = max(maxSum, windowSum)
			windowSum -= arr[windowStart]
			windowStart++
		}
	}

	return maxSum
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func slidingWindow(str string, k int) {
	windowSize := k
	windowStart := 0
	windowEnd := 0

	for windowEnd < len(str) {
		if windowEnd-windowStart+1 == windowSize {
			fmt.Println(str[windowStart : windowEnd+1])
			windowStart++
		}
		windowEnd++
	}
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

// ArrayDiff
func Match(desired int, collateArr []int) bool {
	for _, collate := range collateArr {
		if collate == desired {
			return true
		}
	}
	return false
}
func ArrayDiff(a, b []int) (different []int) {
	for _, desired := range a {
		if !Match(desired, b) {
			different = append(different, desired)
		}
	}
	return different
}

func Digitize(n int) (a []int) {
	if n == 0 {
		a = append(a, 0)
	}
	for ; n > 0; n /= 10 {
		a = append(a, n%10)
	}
	return
}

func Digi(num int) (rev []int) {
	slice := make([]int, 0)

	for _, digit := range strconv.Itoa(num) {
		i, _ := strconv.Atoi(string(digit))
		slice = append(slice, i)
	}
	return slice
}
func ReverseDig(num int) []int {
	str := fmt.Sprintf("%d", num)
	digits := strings.Split(str, "")
	integers := make([]int, len(digits))
	for i, d := range digits {
		n, _ := strconv.Atoi(d)
		integers[len(digits)-i-1] = n
	}
	return integers
}

func main() {
	intarr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	k := 4
	maxSum := maxSumSubarray(intarr, k)
	fmt.Printf("Arr: %#v with k = %d maxSum: %d \n", intarr, k, maxSum)

	sampleString := "abcdefghijklmnopqrstuvwxyz"
	slidingWindow(sampleString, 5)

	var arr = []string{"zone", "abigail", "theta", "form", "libe", "zas"}
	fmt.Printf("LongestConsec: %v\n", LongestConsec(arr, 2))

	var arr2 = []string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"}
	fmt.Printf("LongestConsec: %v\n", LongestConsec(arr2, 1))

	intarr2 := []int{1, 3, 4, 5, 6, 7, 8, 9, 33}
	fmt.Printf("\narr1: %v\n", intarr)
	fmt.Printf("arr2: %v\n", intarr2)
	fmt.Printf("ArrayDiff(arr1, arr2), result: %v. \n", ArrayDiff(intarr, intarr2))
	fmt.Printf("ArrayDiff(arr2, arr1), result: %v. \n", ArrayDiff(intarr2, intarr))

	fmt.Printf("Digitize(35231) result: %v. \n", Digitize(35231))
	fmt.Printf("ReverseDig(35231) result: %v. \n", ReverseDig(35231))
}
