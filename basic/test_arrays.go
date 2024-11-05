package main

import "fmt"

func main() {
	my_array := [6]int{1, 3, 5, 7, 9, 11}
	slice357 := my_array[1:4]
	my_slice := my_array[:]
	fmt.Printf("copy array to slice: %v. \nslice[1:4] index 1 to index 3 %v\n", my_slice, slice357)
	// Testing variadic function. tot()
	fmt.Printf("Total of my_array %v: %d\n\n", my_array, tot(my_array[:]...))

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

	// testing getCostByDay()
	cost := [][]float64{
		{0, 2.5},
		{1, 2.1},
		{1, 2.3},
		{2, 4.5},
		{6, 6.7},
	}
	pl("Cost: %v\n", cost)
	total_cost := getCostByDay(cost)
	pl("total_cost[] %v\n", total_cost)

	for i := 0; i < len(total_cost); i++ {
		pl("Day %d: %f", i+1, total_cost[i])
	}

	// createMatrix()
	fmt.Println("\n--------------\ncreateMatrix(10, 10)\n--------------")

	matrix10_10 := createMatrix(10, 10)
	for i := 0; i < len(matrix10_10); i++ {
		pl("row %d: %v", i, matrix10_10[i])
	}

	// square_matrix()
	fmt.Println("\n--------------\nsquareMatrix(10, 10)\n--------------")

	matrix10 := squareMatrix(matrix10_10)
	for i := 0; i < len(matrix10); i++ {
		pl("row %d: %v", i, matrix10[i])
	}

}

// cost is a 2D array, return sum[] cost of day
func getCostByDay(cost [][]float64) []float64 {
	// get the last value of cost[value][]
	days := cost[len(cost)-1][0] + 1
	// sum[days]
	sum := make([]float64, int(days))

	for i := 0; i < int(days); i++ {
		for j := 0; j < len(cost); j++ {
			if int(cost[j][0]) == i {
				sum[i] += cost[j][1]
			}
		}
	}
	return sum
}

// variadic ...
func tot(num ...int) int {
	total := 0
	for i := 0; i < len(num); i++ {
		total += num[i]
	}
	return total
}

// func pl(s string, value ...interface{}) {
//	 fmt.Println("\n" + s + fmt.Sprintf(s, value...))
// }

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

func createMatrix(row, col int) [][]int {
	matrix := make([][]int, row)
	for i := 0; i < row; i++ {
		matrix[i] = make([]int, col)
	}
	return matrix
}

func squareMatrix(matrix [][]int) [][]int {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			matrix[i][j] = i * j
		}
	}
	return matrix
}
