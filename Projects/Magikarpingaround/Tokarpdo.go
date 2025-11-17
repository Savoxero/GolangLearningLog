package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var magikarp string = `
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⣶⣿⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⣾⠟⠁⢻⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⣾⠟⠁⠀⠀⢸⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⡿⣦⣄⠀⣠⡾⠟⠁⠀⠀⠀⠀⢸⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢹⣇⠈⠛⢿⣟⠁⠀⠀⠀⠀⠀⠀⢸⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⣿⠀⠀⠀⠙⠷⠄⠀⠀⠀⠀⠀⠘⣿⢶⣶⣦⣤⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣿⡆⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠃⠀⠀⢈⣿⠇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣀⣀⣾⡟⠛⢿⣷⣦⣤⣄⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⡿⠋⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⣀⣤⣶⠿⠟⠛⠉⠉⠀⠀⠈⣇⠈⠉⠛⢿⠿⣶⣤⡀⠀⠀⠀⢀⣾⠟⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⢀⣴⡾⠟⠉⠀⠀⠀⠀⠀⠀⠀⠀⠀⠸⣄⣀⣠⠞⠀⠀⠙⠻⣷⣤⣴⠿⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⢀⣴⡿⠋⠀⠀⢀⣠⠤⠤⠤⣄⠀⠀⠀⠀⠀⢇⠀⠈⠳⣄⠀⠀⠀⠈⠙⢿⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣤⡶⠶⣿⣿⡿⣿⠟
⠀⣀⣴⣿⠋⠀⠀⠀⢰⠏⠀⠀⠀⠀⠈⢳⡀⠀⠀⠀⢸⠀⠀⢀⠎⠀⠀⠀⠀⠀⠈⡿⣷⡄⠀⠀⠀⠀⣀⣴⠟⢋⣡⡶⠛⠉⠀⢰⡏⠀
⢸⣟⠙⢧⡀⠀⠀⠀⣟⠀⠀⠀⠃⠀⠀⢰⠇⠀⠀⠀⠀⡗⠒⠣⡀⠀⠀⠀⠀⠀⠀⢱⠈⢿⣆⠀⢀⡼⠋⢀⡴⠟⠁⠀⠀⠀⠀⡿⠀⠀
⠀⠻⣦⡄⠘⡆⠀⠀⠻⣦⡀⠀⠀⣀⡴⠋⠀⠀⢀⣠⡴⠿⠶⠶⠿⢦⣤⣤⣤⣤⣀⣸⣄⠀⣻⣦⠟⠀⣴⣏⠤⠄⠒⠒⠒⠒⣾⠃⠀⠀
⠀⠀⢿⠺⠀⢱⠀⠀⠀⠈⠛⠛⠛⠉⠀⣀⡴⠞⠋⢁⠴⠒⠒⠊⣉⡉⠓⠒⠒⠒⣶⣿⡿⠛⠉⣿⢀⡾⠋⠀⠀⠀⠀⠀⠀⣰⠏⠀⠀⠀
⠀⠀⢸⣀⡇⢸⠀⠀⠀⠀⣠⣤⣀⠀⢸⠋⠀⢀⣴⠥⠒⠊⠉⠁⠀⠀⠀⢀⣴⠟⠉⢠⠃⠀⠀⠈⣾⠁⠀⠀⠀⠀⠀⠀⣴⠃⠀⠀⠀⠀
⠀⣠⣾⣫⠃⡜⠀⠀⠀⠀⠈⠳⣌⢳⡀⠀⢰⣟⠓⠢⠄⡀⠀⠀⠀⢀⣴⠟⠁⠀⠀⡸⠀⠀⠀⣸⡇⠀⠀⠀⠀⠀⠀⣸⠃⠀⠀⠀⠀⠀
⢰⢧⡟⠁⢠⣇⡀⠀⠀⠀⠀⠀⠸⡄⢷⠀⠀⠈⠛⢦⣀⠀⠀⠀⢠⡟⠁⠀⠀⠀⢀⠇⠀⢀⠜⣿⢧⠀⠀⠀⠀⠀⠀⣿⠀⠀⠀⠀⠀⠀
⡞⢸⠛⠒⠿⣷⣬⡓⠢⠤⣀⡀⠀⣇⠸⡆⠀⠀⠀⡾⠛⢷⣄⣰⡟⠀⠀⠀⠀⠀⣞⣠⡴⡟⠀⣿⠈⢆⠀⠀⠀⠀⠀⣿⠀⠀⠀⠀⠀⠀
⢧⠸⡇⠀⠀⠈⠙⠿⣷⣶⣤⣈⡉⣿⠀⡧⠤⣀⣞⠁⠀⣀⠟⠟⠀⠀⠀⣀⣤⣾⠿⠋⠀⣿⠀⣿⠀⠈⠣⡀⠀⠀⢰⡏⠀⠀⠀⠀⠀⠀
⠘⣆⢳⠀⠀⠀⠀⠀⠀⠈⠙⠛⠿⡇⠀⣷⣶⣿⣤⣭⣭⣤⣴⣶⣶⡾⠿⠛⠉⠁⠀⠀⠀⠸⣇⢸⡇⠀⠀⠈⠢⣀⡾⠁⠀⠀⠀⠀⠀⠀
⠀⠹⣦⢧⠀⠀⠀⠀⠀⠀⠀⠀⠀⡇⢸⡿⠃⠀⠀⠀⠀⠀⠀⠀⢻⣦⠀⠀⠀⠀⠀⠀⠀⠀⠹⣦⢿⡄⠀⠀⣴⠟⠁⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠘⣇⢧⠀⠀⠀⠀⠀⠀⠀⢰⠃⢸⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠹⣦⠀⠀⠀⠀⠀⠀⠀⠀⠈⢷⣿⣦⣸⠏⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠸⡎⢧⠀⠀⠀⠀⠀⢀⣾⡄⢸⠀⠀⠀⡀⠀⠀⠀⠀⠀⠀⠀⢹⣆⠀⠀⠀⠀⠀⠀⠀⠀⠀⠙⠻⢿⣦⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⢱⠸⡆⠀⠀⠀⠀⠘⠻⡇⢸⣤⡶⠟⢿⣤⡀⠀⣾⠿⣶⣤⣀⣻⡆⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠈⣤⡇⠀⠀⠀⠀⠀⠀⢻⡘⣇⠀⠀⠀⠙⠻⣶⡟⠀⠀⠉⠙⠛⠟⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⣿⡇⠀⠀⠀⠀⠀⠀⠈⢧⡙⢦⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⢰⡿⠁⠀⠀⠀⠀⠀⠀⠀⠀⠑⢦⣽⣦⣄⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
`
var welcomeMessage string = `
=========================================================================================================
                                   WELCOME TO **TOKARPDO**!
                      A reliable to-do app for your daily use.
                           Enter **help** for the list of commands.
=========================================================================================================`
var tasks []Task
var input string
var a int

type Task struct {
	ID          int
	Description string
	Priority    string
	Completed   bool
	Date        string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(magikarp, welcomeMessage)
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
				ID:          idgenerator(a),
				Description: description,
				Priority:    "medium",
				Completed:   false,
				Date:        time.Now().Format("2006-01-02 15:04:05"),
			}
			tasks = append(tasks, newTask)
			fmt.Println(newTask)

		case "remove":
			id, err := strconv.Atoi(argtokens)
			if err != nil {
				fmt.Println("no argument as ID detected. Enter an ID")
				continue
			}
			if len(tasks) == 0 {
				fmt.Println("no saved tasks currently")
				continue
			}
			found := false /* bool flag for checking whether the task id is inside the slice, if false it means it's not, the moment the loop iterates over a match it's true and as intended*/
			for IDholder, IDmatcher := range tasks {
				fmt.Println("searching for tasks:", tasks[IDholder].ID, tasks[IDholder].Description)
				fmt.Println("comparing IDS:", IDmatcher.ID, id)
				if id == IDmatcher.ID {
					found = true
					fmt.Println("deleted the following task: ID:", tasks[IDholder].ID, "desc:", tasks[IDholder].Description) /* has to be printed before due to the fact that IDholder.ID(the slice number)+1 counteracts the print message, i have tried alternatives but to no avail. */
					tasks = append(tasks[:IDholder], tasks[IDholder+1:]...)
					break
				}

			}
			if found == false {
				fmt.Println("the task with the ID of ", id, " could not be found")
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
func idgenerator(int) int {
	i := 0
	for i = 0; i <= 1000; i++ {
		if i > len(tasks) {
			break
		} else {
			i = len(tasks) + 1
		}

	}
	return i
}
