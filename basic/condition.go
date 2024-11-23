package main

import (
	"fmt"
)

func main() {

	goColor.color.Pl("This is numbers ( 0123456789 ).  %d", 9876543210)

	printPrimes(20)
}

// Prime = only devided by 1 or it self
func printPrimes(max int) {
	fmt.Printf("Print prime numbers up to %d\n", max)
	for n := 2; n < max+1; n++ {
		if n == 2 {
			fmt.Println(n) // 2 is prime ? even nr ?
			continue
		}
		if n%2 == 0 { // if even, not prime
			continue
		}
		isPrime := true
		for i := 3; i*i < n+1; i++ {
			if n%i == 0 {
				isPrime = false
				break
			}
		}

		if !isPrime {
			continue
		}
		fmt.Println(n)
	}
}
