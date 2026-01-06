package CMD

import (
	"errors"
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

func Add(argtokens string) ([]Task, error) {
	var taskprint []Task
	argtokens = strings.TrimSpace(argtokens)

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
		taskprint = append(taskprint, newTask)
	}
	if len(taskprint) == 0 {
		return taskprint, errors.New("Empty input detected, please enter a valid description")
	} else {
		return taskprint, nil

	}
}

func IDgenerator(a int) int {
	a++
	return a
}
