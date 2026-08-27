package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	fileName     = "todo_list.txt"
	fileNameTemp = "todo_list_temp.txt"
	prefix       = "- "
	prefixLen    = len(prefix)
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	todos, err := loadAndParseTodos()
	if err != nil {
		fmt.Println("Error opening the file with the list of tasks", err)
		return
	}

	for {
		clearTerminal()

		fmt.Println("(1) Add a task to the list")
		fmt.Println("(2) Show the list of tasks")
		fmt.Println("(3) Edit a task")
		fmt.Println("(4) Delete a task")
		fmt.Println("(5) Delete all tasks")
		fmt.Println("(0) Exit")

		fmt.Println("Select an option:")

		input := readTrimmedString(reader)

		clearTerminal()

		switch strings.TrimSpace(input) {
		case "1":
			todos = addTodo(reader, todos)
		case "2":
			printTodos(reader, todos, true)
		case "3":
			todos = editTodo(reader, todos)
		case "4":
			todos = deleteTodo(reader, todos)
		case "5":
			todos = deleteAll(todos)
		case "0":
			return
		}
	}
}

func loadAndParseTodos() ([]string, error) {
	todos := []string{}

	file, err := os.Open(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return todos, nil
		}
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, prefix) {
			todos = append(todos, line[prefixLen:])
		} else if line != "" {
			continue
		}
	}

	return todos, scanner.Err()
}

func printTodos(reader *bufio.Reader, todos []string, withWait bool) {
	for i, todo := range todos {
		fmt.Println(strconv.Itoa(i+1) + ". " + todo)
	}

	if len(todos) == 0 {
		fmt.Println("The list of tasks is empty")
	}

	if withWait {
		fmt.Print("\nPress \"Enter\" to continue")
		reader.ReadString('\n')
	}
}

func addTodo(reader *bufio.Reader, todos []string) []string {
	printTodos(reader, todos, false)

	fmt.Print("\nEnter the text of the task: ")
	input := readTrimmedString(reader)
	if input == "" {
		return todos
	}

	todos = append(todos, input)

	err := saveTodos(todos)
	if err != nil {
		todos = todos[:len(todos)-1]
	}

	return todos
}

func deleteTodo(reader *bufio.Reader, todos []string) []string {
	printTodos(reader, todos, false)

	fmt.Print("\nEnter the number of the task to delete: ")
	num, err := readInputTodoNumber(reader, todos)
	if err != nil {
		return todos
	}

	idx := num - 1
	oldTodo := todos[idx]
	todos = append((todos)[:idx], (todos)[idx+1:]...)

	err = saveTodos(todos)
	if err != nil {
		reinserted := append([]string{oldTodo}, todos[idx:]...)
		todos = append(todos[:idx], reinserted...)
	}

	return todos
}

func deleteAll(todos []string) []string {
	emptyTodos := []string{}

	if err := saveTodos(emptyTodos); err != nil {
		return todos
	}

	return emptyTodos
}

func editTodo(reader *bufio.Reader, todos []string) []string {
	printTodos(reader, todos, false)

	fmt.Print("\nEnter the number of the task to edit: ")
	num, err := readInputTodoNumber(reader, todos)
	if err != nil {
		return todos
	}

	idx := num - 1

	fmt.Print("\nEnter the text of the task: ")

	input := readTrimmedString(reader)

	if input == "" {
		return todos
	}

	oldTodo := todos[idx]
	todos[idx] = input

	err = saveTodos(todos)
	if err != nil {
		todos[idx] = oldTodo

		return todos
	}

	return todos
}

func saveTodos(todos []string) error {
	f, err := os.Create(fileNameTemp)
	if err != nil {
		printErrorWithWait("Error saving the list of tasks. Changes were not saved")

		return err
	}

	for _, todo := range todos {
		_, err := f.WriteString(prefix + todo + "\n")
		if err != nil {
			f.Close()
			removeTempFile()
			printErrorWithWait("Error saving the list of tasks. Changes were not saved")

			return err
		}
	}

	if err := f.Close(); err != nil {
		removeTempFile()
		printErrorWithWait("Error closing temporary file")
		return err
	}

	err = os.Rename(fileNameTemp, fileName)
	if err != nil {
		removeTempFile()
		printErrorWithWait("Error renaming temporary file")

		return err
	}

	return nil
}

func removeTempFile() {
	// If deleting the temporary file did not work, it is not a critical error
	_ = os.Remove(fileNameTemp)
}

func readInputTodoNumber(reader *bufio.Reader, todos []string) (int, error) {
	input := readTrimmedString(reader)

	if input == "" {
		return 0, fmt.Errorf("Empty input")
	}

	num, err := strconv.Atoi(input)

	if err != nil {
		printErrorWithWait("Entered number is incorrect!")

		return 0, err
	}

	if num < 1 {
		printErrorWithWait("Entered number should be greater than 0!")
		return 0, fmt.Errorf("Wrong number")
	}

	if num > len(todos) {
		printErrorWithWait("There is no such number in the list of tasks!")
		return 0, fmt.Errorf("Wrong number")
	}

	return num, nil
}

func readTrimmedString(reader *bufio.Reader) string {
	input, err := reader.ReadString('\n')

	if err != nil {
		printErrorWithWait("Error reading data")
		return ""
	}

	input = strings.TrimSpace(input)

	return input
}

func printErrorWithWait(text string) {
	fmt.Println(text)
	time.Sleep(3 * time.Second)
}

func clearTerminal() {
	fmt.Print("\033[H\033[2J\033[3J")
}
