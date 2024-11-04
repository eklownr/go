package main

import (
	"fmt"
)

type cost struct {
	day   int
	price float64
}

func main() {
	person := struct {
		name string
		age  int
	}{
		name: "John",
		age:  30,
	}
	fmt.Println(person)

	type car struct {
		make  string
		model string
		wheel struct {
			radius int
		}
	}

	type truck struct {
		car
		bedSize int
	}

	truck1 := truck{
		car: car{
			make:  "Ford",
			model: "F-150",
			wheel: struct {
				radius int
			}{
				radius: 18,
			},
		},
		bedSize: 180,
	}
	fmt.Printf("Truck1 make and model: %s %s.\nWith bedsize: %d cm\n",
		truck1.make, truck1.model, truck1.bedSize)

	//////////////////////////
	// testing getCostByDay()
	test_list := []cost{
		{0, 2.5},
		{1, 2.1},
		{1, 2.3},
		{2, 4.5},
		{6, 6.7},
	}
	fmt.Printf("Cost: %v\n", test_list)
	total_cost := get_CostByDay(test_list)
	fmt.Printf("total_cost[] %v\n", total_cost)

	for i := 0; i < len(total_cost); i++ {
		fmt.Printf("Day %d: %f\n", i+1, total_cost[i])
	}

}

func get_CostByDay(cost []cost) []float64 {
	days := cost[len(cost)-1].day + 1
	sum := make([]float64, int(days))
	for _, c := range cost {
		sum[c.day] += c.price
	}
	return sum
}
