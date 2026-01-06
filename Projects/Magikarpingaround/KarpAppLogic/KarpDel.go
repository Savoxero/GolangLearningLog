package CMD

import (
	"errors"
	"strconv"
	"strings"
)

var Taskdeletion []Task

func Del(argtokens string) error {
	multdel := strings.Split(argtokens, "&")
	for _, IDIN := range multdel {
		found := false
		IDIN = strings.TrimSpace(IDIN)
		if IDIN == "" {
			return errors.New("Empty input detected. Enter a valid ID")
		}
		if len(tasks) == 0 {
			return errors.New("No tasks currently saved.")
		}
		id, err := strconv.Atoi(IDIN)
		if err != nil {
			return errors.New("Couldn't verify the ID; Please enter a valid ID")
		}

		for i, d := range tasks {
			if id == d.ID {
				found = true
				Taskdeletion = append(Taskdeletion, tasks[i])
				tasks = append(tasks[:i], tasks[i+1:]...)
				break

			}

		}
		if !found {
			return errors.New("The number with the ID of " + IDIN + " Has not been found")
		}
	}

	return nil

}
