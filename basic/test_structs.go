package main

import (
	"fmt"
)

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

}
