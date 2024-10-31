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

	check_for_error("go Golang go")
	check_for_error("") // Hello("") will return an error
}

func check_for_error(s string) {
	message, err := greetings.Hello(s)
	// If an error was returned, print it to the console and exit the program.
	if err != nil {
		log.Fatal(err)
	}
	// If no error was returned, print the returned message
	fmt.Println(message)
}
