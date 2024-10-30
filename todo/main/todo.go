package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	//"github.com/fatih/color"
)

func main() {
	task := []string{}        // emty string array
	task = readTodoFile(task) // add all lines from file todo.txt

	welcom := ("******* Welcome to my Todo list app *******")
	color.Blue(welcom)
	printTask(task)

	reader := bufio.NewReader(os.Stdin)
	for {
		color.Cyan("Enter command (q, a, d, p): ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		text = strings.TrimSpace(text)
		text = strings.ToLower(text)
		if text == "q" {
			color.Green("Bye - Exit todo app - try to save file")
			writeToTodoFile(task, "todo.txt")
			break
		} else if text == "d" {
			color.Red("Delete task:")
			fmt.Print("Enter task number: ")
			var index int
			fmt.Scanln(&index)
			task = delTask(task, index-1)
			printTask(task)
		} else if text == "a" {
			fmt.Println("Add task:")
			fmt.Print("Enter task: ")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan() // read a whole line
			line := scanner.Text()
			task = addTask(task, line)
			printTask(task)
		} else if text == "p" {
			printTask(task)
		}
		fmt.Println("\n'd' delete - 'a' add task - 'q' quit - 'p' print list:")
	}
}

// open file todo.txt and return all lines as []string
func readTodoFile(taskList []string) []string {
	lines, err := openFile("todo.txt")
	if err != nil {
		log.Fatal(err)
	}
	//for _, line := range lines {
	//fmt.Println(line)
	taskList = append(taskList, lines...)
	//}
	return taskList
}

func printTask(l []string) {
	color.Yellow("\nList of my Todos:")
	for index, item := range l {
		fmt.Printf("%d: %s\n", index+1, item)
	}
}

// Delete task
func delTask(taskList []string, index int) []string {
	fmt.Println("\n****** Delete task: ", index+1, "*******")
	buf := []string{}
	for i, taskValue := range taskList {
		if i != index {
			buf = append(buf, taskValue)
		} else {
			fmt.Printf("Task %d: '%s' deleted \n", index+1, taskValue)
		}
	}
	return buf
}

// add item to task and update tasklist
func addTask(task []string, s string) []string {
	fmt.Printf("\n****** Add task: %s *******\n", s)
	task = append(task, s)
	return task
}

// fmt.println....
func pl(s string) {
	fmt.Println(s)
}

func openFile(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func saveFile(filename string, NewTaskList []string) error {
	// Open the file in write mode
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	// Join the task list strings with newline characters
	taskListStr := strings.Join(NewTaskList, "\n")

	// Write the task list string to the file
	_, err = f.WriteString(taskListStr)
	if err != nil {
		return err
	}

	return nil
}

func writeToTodoFile(newTaskList []string, filename string) {
	err := saveFile(filename, newTaskList)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("File saved successfully!")
	}
}
