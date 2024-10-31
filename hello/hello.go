package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including the log entry prefix
	//  and a flag '0' to disable printing the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(11)

	multi_greetings()
	hello_name("Golang")
	//hello_name("") // Hello("") will return an error

	// receiving map
	my_map := get_name_map()
	fmt.Println(my_map)
	// print key and value
	for k, v := range my_map {
		fmt.Printf("key: %v. value: %v \n", k, v)
	}
}

func multi_greetings() {
	// A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin"}

	// Request greeting messages for the names.
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	// If no error was returned, print the returned map of
	// messages to the console.
	fmt.Println(messages)
}

func get_name_map() map[string]string {
	names := []string{"Gladys", "Samantha", "Darrin"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	return messages

}

func hello_name(name string) {
	message, err := greetings.Hello(name)
	// If an error was returned, print it to the console and exit the program.
	if err != nil {
		log.Fatal(err)
	}
	// If no error was returned, print the returned message
	fmt.Println(message)
}
