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
			if err != nil {
				fmt.Println("invalid", id, "Please enter a valid ID")
				continue
			}
			for IDfinder := range tasks {
				if tasks[IDfinder].ID == id {
					tasks = append(tasks[:IDfinder], tasks[IDfinder+1:]...)
					fmt.Println("Deleted Task:", id)
					break
				} else if len(tasks) == 0 {
					fmt.Println(id, "no current tasks logged")
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
