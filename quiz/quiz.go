package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFileName := flag.String("csv", "problems.cvs", "A csv file in format of 'qestion, answer'")
	flag.Parse()

	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open CSV file: %s\n", *csvFileName))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provied csv file.")
	}

	problems := parseLines(lines)
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.q)
		var answer string
		fmt.Scanf("%s", &answer)
		if answer == p.a {
			pl("Correct answer!")
		}
	}

}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: line[1],
		}
	}
	return ret
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func pl(s string) {
	fmt.Println(s)
}
