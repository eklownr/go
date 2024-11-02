package main

import "fmt"

func main() {
	my_array := [6]int{1, 3, 5, 7, 9, 11}
	slice357 := my_array[1:4]
	my_slice := my_array[:]
	fmt.Printf("copy array to slice: %v. slice[1:4] index 1 to index 3 %v", my_slice, slice357)

	slice_from_make := make([]int, 5)
	fmt.Println(slice_from_make)
}
