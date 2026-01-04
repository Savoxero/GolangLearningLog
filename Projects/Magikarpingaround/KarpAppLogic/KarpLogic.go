package CMD

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Task struct {
	ID          int
	Description string
	Priority    string
	Completion  bool
	Date        string
}

var tasks []Task
var RiseID int

func Add(argtokens string) (Task, error) {
	var taskprint Task
	argtokens = strings.TrimSpace(argtokens)
	/// edge case for empty input necessary fix, {0 false}
	if argtokens == "" {
		return taskprint, errors.New("Empty input detected")
	}
	descriptions := strings.Split(argtokens, "&")
	for _, desc := range descriptions {
		desc = strings.TrimSpace(desc)
		if desc == "" {
			continue
		}
		RiseID = IDgenerator(RiseID)
		newTask := Task{
			ID:          RiseID,
			Description: desc,
			Priority:    "medium",
			Completion:  false,
			Date:        time.Now().Format("2006-01-02 15:04:05"),
		}
		tasks = append(tasks, newTask)
		taskprint = newTask
	}
	return taskprint, nil
}

func Del(argtokens string) {
	multdel := strings.Split(argtokens, "&")
	for _, IDIN := range multdel {
		found := false
		IDIN = strings.TrimSpace(IDIN)
		if IDIN == "" {
			fmt.Println("No id detected/error")
			continue
		}
		if len(tasks) == 0 {
			fmt.Println("no saved tasks currently")
			break

		}
		id, err := strconv.Atoi(IDIN)
		if err != nil {
			fmt.Println(" Error Occured, Invalid Input, an ID is necessary")
			continue
		}

		for i, d := range tasks {
			if id == d.ID {
				found = true
				fmt.Println("deleted the following task: ID:", d.ID, "desc:", d.Description) /* has to be printed before due to the fact that i(the slice number)+1 counteracts the print message, i have tried alternatives but to no avail. */
				tasks = append(tasks[:i], tasks[i+1:]...)
				continue
			}

		}
		if found == false {
			fmt.Println("the task with the ID of", id, " could not be found")
		}

	}

}
func Print(id string) {
	argtokens := strings.TrimSpace(id)
	multiput := strings.Split(argtokens, "&")
	if len(tasks) == 0 && argtokens == "" {
		fmt.Println("no saved tasks currently")

	}
	if argtokens != "" {
		for _, stRang := range multiput {
			boolflag := false
			stRang = strings.TrimSpace(stRang)
			usrInerr := stRang
			dada, err := strconv.Atoi(stRang)
			if err != nil {
				fmt.Println(usrInerr, "Is an invalid argument. Please Enter an ID or Type Check With No ID")
				continue
			}
			for i := range tasks {
				if tasks[i].ID == dada {
					fmt.Println("Task Number:", tasks[i].ID, "// Desc:", tasks[i].Description, "// Time added/Created:", tasks[i].Date, "// Priority:", tasks[i].Priority, "// Completion:", tasks[i].Completion)
					boolflag = true
					continue
				}

			}
			if boolflag == false {
				fmt.Println("The task with the ID of", dada, "could not be found.")
				continue
			}
		}
	}
	if argtokens == "" {
		fmt.Println(tasks)

	}
}
func IDgenerator(a int) int {
	a++
	return a
}
