package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var input string

var tasks []Task

type Task struct {
	ID          int
	Description string
	Priority    string
	Completed   bool
	Date        string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to clotodo, a to do app for your daily use. Enter help for the list of commands")
	for {
		fmt.Print("> ")
		scanner.Scan()
		input = scanner.Text()
		filter := strings.Fields(input)
		cmtokens := filter[0]
		argtokens := strings.Join(filter[1:], " ")
		input = cmtokens
		switch input {
		case "help":
			fmt.Println("")

		case "add":
			description := argtokens
			newTask := Task{
				ID:          len(tasks) + 1,
				Description: description,
				Priority:    "medium",
				Completed:   false,
				Date:        time.Now().Format("2006-01-02 15:04:05"),
			}
			fmt.Println(newTask)
			tasks = append(tasks, newTask)
		case "remove":
			id, err := strconv.Atoi(argtokens)
			for taskfinder := range tasks {
				if tasks[taskfinder].ID == id {
					break
				} else if tasks[taskfinder].ID != id {
					fmt.Println("invalid", err, " number, please enter a valid Task ID")
					continue
				}

			}

		case "modify":
		case "default":
		case "exit":
			fmt.Println("Goodbye, come again")
			return
		}
	}
}
