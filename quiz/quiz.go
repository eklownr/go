package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFileName := flag.String("csv", "problems.cvs", "A csv file in format of 'qestion, answer'")
	timeLimit := flag.Int("limit", 30, "The time limit for the quiz in seconds.")
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

	correct := 0
	problems := parseLines(lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.q)

		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s", &answer) // read from standad in
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nYou reached timelilmit, you scored %d out of %d. \n", correct, len(problems))
			return // Stop program
		case answer := <-answerCh:
			if answer == p.a {
				pl("Correct answer!")
				correct++
			}
		}
	}

	fmt.Printf("You scored %d out of %d. \n", correct, len(problems))

}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
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
