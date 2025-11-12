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
		if len(filter) <= 0 {
			fmt.Println("no input detected, enter help for a list of commands")
			continue
		}
		cmtokens := filter[0]
		argtokens := strings.Join(filter[1:], " ")
		input = cmtokens
		switch input {
		case "help":
			fmt.Println("")

		case "add":
			description := argtokens
			if description == "" {
				fmt.Println("invalid argument")
				break
			}
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
			if err != nil {
				fmt.Println("no argument as ID detected. Enter an ID")
				continue
			}
			if len(tasks) == 0 {
				fmt.Println("no task of this id has been found")
				continue
			}
			for IDfinder := range tasks {
				if id > tasks[IDfinder].ID {
					fmt.Println("no task of this id has been found")
					break
				}
				if id == tasks[IDfinder].ID {
					tasks = append(tasks[:IDfinder], tasks[IDfinder+1:]...)
					fmt.Println("Deleted Task:", tasks[IDfinder].ID, tasks[IDfinder].Description)
					fmt.Println(tasks[0])
					break
				}
			}

		case "modify":
		case "show":
			fmt.Println("the current saved tasks:", tasks)
		case "default":
		case "exit":
			fmt.Println("Goodbye, have a good day/night")
			return

		}
	}

}
