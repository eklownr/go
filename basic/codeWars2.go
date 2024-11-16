package main

import (
	"fmt"
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

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	k := 4
	maxSum := maxSumSubarray(arr, k)
	fmt.Printf("Arr: %#v with k = %d maxSum: %d \n", arr, k, maxSum)

	sampleString := "abcdefghijklmnopqrstuvwxyz"
	slidingWindow(sampleString, 5)
}
