package Domain

import (
	"errors"
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

func AddTask(input []string) ([]Task, error) {
	var tempholder []Task
	var errs []error
	tempholder = nil

	if len(input) == 0 {
		return nil, errors.New("no input detected")
	}
	for desci, desc := range input {
		if desc == "" {
			if len(input) > 1 {
				errs = append(errs, errors.New("Empty description, please enter a valid description."))
				continue
			} else {
				return nil, errors.New("Empty description, please enter a valid description.")
			}
		}
		RiseID++
		Newtask := Task{
			ID:          RiseID,
			Description: input[desci],
			Priority:    "Medium",
			Completion:  false,
			Date:        time.Now().Format("2006-01-02 15:04:05")}
		tasks = append(tasks, Newtask)
		tempholder = append(tempholder, Newtask)
	}
	if len(errs) > 0 {
		return tempholder, errors.New("╔════════════════════════════════════════════════════════════════════╗\n║  ⚠️  Some empty input detected!⚠️                                    ║\n║  Use: add: Take out the doggy & clean the kitchen                  ║\n║  Tasks with valid input still saved.                               ║\n╚════════════════════════════════════════════════════════════════════╝")
	}
	return tempholder, nil
}
func DeleteTask(input []int) ([]Task, error) {
	var errs []error
	if len(tasks) == 0 {
		return nil, errors.New("No saved tasks currently")
	}
	found := false
	for _, id := range input {
		for d, task := range tasks {
			if id == task.ID {
				tasks = append(tasks[:d], tasks[d+1:]...)
				found = true
				break
			}
		}
		if !found {
			errs = append(errs, errors.New("The ID with the number of: "+strconv.Itoa(id)+" has not been found"))
		}
	}
	if len(errs) > 0 {
		var msgs []string
		for _, i := range errs {
			msgs = append(msgs, i.Error())
			return tasks, errors.New(strings.Join(msgs, "; "))
		}
	}
	return tasks, nil
}
func KarpPrint(input []int) ([]Task, error) {
	var errs []error
	var Taskprintholder []Task
	Taskprintholder = nil
	if len(tasks) == 0 {
		return nil, errors.New("No saved tasks currently")
	}
	if input == nil {
		return tasks, nil
	}
	found := false
	for _, id := range input {
		for i, task := range tasks {
			if id == task.ID {
				Taskprintholder = append(Taskprintholder, tasks[i])
				found = true
			}
		}
		if !found {
			errs = append(errs, errors.New("The ID with the number of: "+strconv.Itoa(id)+" Is an invalid argument. Please Enter a valid ID or Type print With No ID"))
		}
		if len(errs) > 0 {
			var msgs []string
			for _, i := range errs {
				msgs = append(msgs, i.Error())
				return Taskprintholder, errors.New(strings.Join(msgs, "; "))
			}
		}
	}
	return Taskprintholder, nil
}
