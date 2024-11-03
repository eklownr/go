package main

import "fmt"

func main() {
	my_array := [6]int{1, 3, 5, 7, 9, 11}
	slice357 := my_array[1:4]
	my_slice := my_array[:]
	fmt.Printf("copy array to slice: %v. \nslice[1:4] index 1 to index 3 %v\n", my_slice, slice357)

	slice_from_make := make([]int, 5)
	fmt.Printf("The type of slice_from_make: %T \nThe value: %v \nCapacity: %v \n\n",
		slice_from_make, slice_from_make, cap(slice_from_make))
	matrix := make([][]int, 5)
	//fmt.Printf("The type of matrix: %T \nThe value: %v \n", matrix, matrix)
	matrix = [][]int{
		{1, 1, 1, 1, 1},
		{2, 2, 2, 2, 2},
		{1234, 3, 3, 3, 3},
		{4, 4, 4, 4, 4},
		{5, 5, 5, 5, 5},
	}
	fmt.Printf("The type of matrix: %T \nThe value: %v \n", matrix, matrix)
	fmt.Printf("matrix[2][0]: %d \n\n", matrix[2][0])

	make_slice := make([]int, 5, 10)
	fmt.Printf("The type of make_slice: %T \nThe value: %v \nthe length: %v \nThe capacity: %v\n",
		make_slice, make_slice, len(make_slice), cap(make_slice))

	fmt.Printf("Total of matrix %v: %d\n\n", my_array, tot(my_array[:]...))

	cost := [][]float64{
		{0, 2.5},
		{1, 2.1},
		{1, 2.3},
		{2, 4.5},
		{6, 6.7},
	}
	total_cost := getCostByDay(0, cost)

	fmt.Printf("Total of cost %v: \n%f\n", cost, total_cost)

	// pl(string, interface). Replace for fmt.Println()
	pl("total %v", total_cost)
	pl("Testing int %v", 1234)
	my_string := "Hello World"
	pl("String %v", my_string)
	pl("Total of cost %v: \n%f\n", cost, total_cost)
}

// cost is a 2D array, return total cost of day
func getCostByDay(day int, cost [][]float64) float64 {
	for i := 0; i < len(cost); i++ {
		if cost[i][0] == float64(day) {
			return cost[i][1]
		}
	}
	return 0
}

// replace for fmt.Printf(string, multiple values)
func pl(s string, value ...interface{}) {
	i := len(value)
	if i == 0 {
		fmt.Println(s)
	}
	if i == 1 {
		fmt.Printf(s, value[0])
		fmt.Println()
	}
	if i == 2 {
		fmt.Printf(s, value[0], value[1])
		fmt.Println()
	}
	if i == 3 {
		fmt.Printf(s, value[0], value[1], value[2])
		fmt.Println()
	}
	if i == 4 {
		fmt.Printf(s, value[0], value[1], value[2], value[3])
		fmt.Println()
	}
}

// variadic ...
func tot(num ...int) int {
	total := 0
	for i := 0; i < len(num); i++ {
		total += num[i]
	}
	return total
}
