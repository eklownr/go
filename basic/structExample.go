package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
}

func main() {

	person := Person{
		firstName: "John",
		lastName:  "Doe",
		age:       30,
	}

	fmt.Println(person) // age is 30
	copy_of_person := changeCopy(person)
	fmt.Println(person) // age is 30
	changeRef(&person)
	fmt.Println(person)         // age is 40
	fmt.Println(copy_of_person) // age is 60
}

func changeCopy(person Person) Person {
	person.age = 60
	return person
}
func changeRef(person *Person) {
	person.age = 40
}
